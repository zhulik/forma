package codegen

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/zhulik/forma/fn"
)

//go:embed overrides/*.override.yaml
var embeddedOverrides embed.FS

// Per-source overrides merge on top of TSX analysis when heuristics miss props,
// events, or compound slots.
//
// Resolution order:
//  1. Sidecar next to the component: <base>.override.yaml (e.g. input.override.yaml)
//  2. Embedded maintainer defaults under overrides/<base>.override.yaml (shipped for
//     shadcn input/checkbox forwardRef sources that expose no Prop interface)
//
// YAML schema:
//
//	kind: leaf          # optional: leaf | compound
//	name: Input         # optional Go/widget name
//	props:
//	  - name: OnChange
//	    goType: "func(ctx ui.Context, value string)"
//	  - name: Value
//	    goType: string
//	slots:              # compound only; replaces analyzed slots when non-empty
//	  - field: Header
//	    exportName: CardHeader
//
// Known goType values: string, bool, float64, func(ctx ui.Context),
// func(ctx ui.Context, value string), func(ctx ui.Context, checked bool).
// Matching prop names replace analyzed fields; new names are appended.

// SourceOverride is the committed YAML schema for one TSX file.
type SourceOverride struct {
	Kind  string         `yaml:"kind"`
	Name  string         `yaml:"name"`
	Props []PropOverride `yaml:"props"`
	Slots []SlotOverride `yaml:"slots"`
}

// PropOverride forces a Go field name and type on the analyzed widget.
type PropOverride struct {
	Name   string `yaml:"name"`
	GoType string `yaml:"goType"`
}

// SlotOverride forces a compound slot field ↔ export mapping.
type SlotOverride struct {
	Field      string `yaml:"field"`
	ExportName string `yaml:"exportName"`
}

func overridePathForSource(sourcePath string) string {
	ext := filepath.Ext(sourcePath)
	base := strings.TrimSuffix(sourcePath, ext)
	return base + ".override.yaml"
}

func loadSourceOverride(sourcePath string) (SourceOverride, bool, error) {
	path := overridePathForSource(sourcePath)
	raw, err := os.ReadFile(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return SourceOverride{}, false, fmt.Errorf("read override %s: %w", path, err)
		}
		fileBase := strings.TrimSuffix(filepath.Base(sourcePath), filepath.Ext(sourcePath))
		raw, err = embeddedOverrides.ReadFile("overrides/" + fileBase + ".override.yaml")
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				return SourceOverride{}, false, nil
			}
			return SourceOverride{}, false, fmt.Errorf("read embedded override %s: %w", fileBase, err)
		}
		path = "embedded:overrides/" + fileBase + ".override.yaml"
	}
	var ov SourceOverride
	if err := yaml.Unmarshal(raw, &ov); err != nil {
		return SourceOverride{}, false, fmt.Errorf("parse override %s: %w", path, err)
	}
	return ov, true, nil
}

func applySourceOverride(sourcePath string, widget Widget) (Widget, error) {
	ov, ok, err := loadSourceOverride(sourcePath)
	if err != nil {
		return widget, err
	}
	if !ok {
		return widget, nil
	}
	return mergeOverride(widget, ov), nil
}

// MergeOverride applies ov on top of widget (exported for unit tests).
func MergeOverride(widget Widget, ov SourceOverride) Widget {
	return mergeOverride(widget, ov)
}

func mergeOverride(widget Widget, ov SourceOverride) Widget {
	if ov.Kind != "" {
		widget.Kind = ov.Kind
	}
	if ov.Name != "" {
		widget.Name = ov.Name
		if widget.Kind == "leaf" {
			widget.ExportName = ov.Name
		} else {
			widget.RootExport = ov.Name
		}
	}
	if len(ov.Props) > 0 {
		widget.Props = mergePropOverrides(widget.Props, ov.Props)
	}
	if len(ov.Slots) > 0 {
		widget.Slots = fn.Map(ov.Slots, func(s SlotOverride) SlotField {
			return SlotField(s)
		})
		slices.SortFunc(widget.Slots, func(a, b SlotField) int {
			return strings.Compare(a.Field, b.Field)
		})
	}
	return widget
}

func mergePropOverrides(props []PropField, overrides []PropOverride) []PropField {
	byName := fn.Associate(props, func(p PropField) (string, PropField) {
		return p.Name, p
	})
	for _, o := range overrides {
		if o.Name == "" || o.GoType == "" {
			continue
		}
		byName[o.Name] = PropField(o)
	}
	return fn.Map(slices.Sorted(maps.Keys(byName)), func(name string) PropField {
		return byName[name]
	})
}

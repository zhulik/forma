package formaconfig

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const FileName = "forma.yaml"

// Load reads forma.yaml from dir.
func Load(dir string) (Project, error) {
	path := filepath.Join(dir, FileName)
	raw, err := os.ReadFile(path)
	if err != nil {
		return Project{}, err
	}
	var p Project
	if err := yaml.Unmarshal(raw, &p); err != nil {
		return Project{}, err
	}
	return p, nil
}

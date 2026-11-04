package codegen

import (
	"context"

	"github.com/zhulik/pal"

	"github.com/zhulik/forma/cmd/forma/internal/core"
)

// Provide registers the codegen runner for pal injection.
func Provide() pal.ServiceDef {
	return pal.Provide[core.CodegenRunner](&adapter{})
}

type adapter struct {
	Runner Runner
}

func (a *adapter) Generate(ctx context.Context, req core.CodegenRequest) error {
	return a.Runner.Generate(ctx, Request{
		ComponentsDir: req.ComponentsDir,
		CodegenDir:    req.CodegenDir,
		Package:       req.Package,
		RegistryOut:   req.RegistryOut,
		FormaUIImport: req.FormaUIImport,
	})
}

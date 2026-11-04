package clierrors

import "errors"

var (
	ErrNotEmpty              = errors.New("directory is not empty")
	ErrConfigExists          = errors.New("forma.yaml already exists")
	ErrModuleRequired        = errors.New("module path is required; pass it as: forma init <module>")
	ErrUnknownTemplate       = errors.New("unknown template")
	ErrNotImplemented        = errors.New("not implemented yet")
	ErrWebPackageMissing     = errors.New("web tree is missing package.json")
	ErrComponentNameRequired = errors.New("component name is required; pass it as: forma add <component>")
)

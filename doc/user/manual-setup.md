# Manual project setup

Use this when you prefer not to run **`forma init`**, or when you only need to add **`forma.yaml`** to an existing Go module.

## Prerequisites

Same toolchain as [Install](install.md): Go matching `go.mod`, plus Node and GoScript for **`forma sync`** / **`forma dev`**.

## 1. Create the Go module

```bash
mkdir myapp && cd myapp
go mod init github.com/you/myapp
```

## 2. Add forma.yaml

Create **`forma.yaml`** at the repo root (or run **`forma init --only-config github.com/you/myapp`** to generate it):

```yaml
module: github.com/you/myapp

ui:
  mount: app-root

web:
  path: web
```

## 3. Minimal main.go

```go
package main

import (
	. "github.com/zhulik/forma/ui" //nolint:revive,staticcheck // DSL
)

func main() {
	UIMount("app-root", App())
}

func App() Component {
	return Container{}.With(
		Heading{Level: 1, Text: "Hello from forma"},
	)
}
```

The **`minimal`** embed template from **`forma init`** matches this layout. Owned widgets (`Button`, `Input`, layout, …) come from **`github.com/zhulik/forma/ui`**.

## 4. Depend on the forma library

Before the first tagged release, point at a local checkout or proxy version:

```bash
go mod edit -require=github.com/zhulik/forma@v0.0.0
# Local development only:
go mod edit -replace=github.com/zhulik/forma=/path/to/forma
go mod tidy
```

After **`forma`** is published, **`forma sync`** will reconcile requires for you.

## 5. Web tree

Copy the scaffold under `cmd/forma/templates/scaffolds/minimal/web/` (or run **`forma init`** into a temp dir and copy), then:

```bash
forma sync
```

## 6. Verify

```bash
go build -o /dev/null .
```

## Related

- [Install](install.md)
- [CLI reference](cli.md)
- [Project layout](project-layout.md)

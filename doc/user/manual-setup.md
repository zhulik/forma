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
  components: []

codegen:
  dir: internal/ui
  package: ui

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

The **`minimal`** embed template from **`forma init`** matches this layout.

## 4. Depend on the forma library

Before the first tagged release, point at a local checkout or proxy version:

```bash
go mod edit -require=github.com/zhulik/forma@v0.0.0
# Local development only:
go mod edit -replace=github.com/zhulik/forma=/path/to/forma
go mod tidy
```

After **`forma`** is published, **`forma sync`** (stub today, full later) will reconcile requires for you.

## 5. Verify

```bash
go build -o /dev/null .
```

Browser runtime and **`internal/ui`** codegen arrive with P3; this step only checks the Go authoring model.

## Related

- [Install](install.md)
- [CLI reference](cli.md)
- [Project layout](project-layout.md)

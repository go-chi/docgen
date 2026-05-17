# docgen

Route documentation generator for [chi](https://github.com/go-chi/chi) routers.

## Install

```bash
go get github.com/go-chi/docgen
```

## Usage

Generate markdown API docs from your router:

```go
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})

	md := docgen.NewMarkdownDoc(r)
	md.Opts.ProjectPath = "github.com/my/app"
	md.Opts.URLMap = map[string]string{
		"github.com/go-chi/chi/v5": "https://github.com/go-chi/chi/tree/master",
	}
	println(md.Markdown())
}
```

See `docgen_test.go` and `raml/raml_test.go` for more examples.

## License

MIT

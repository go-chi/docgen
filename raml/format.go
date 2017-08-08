package raml

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pressly/chi"
)

type Format struct {
	Middleware         bool
	UnexportedHandlers bool
}

type FormatFn func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) (*Resource, error)

func (raml *RAML) AddResourcesFmt(r chi.Routes, fn FormatFn) error {
	return chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		resource, err := fn(method, route, handler, middlewares...)
		if err != nil {
			return err
		}
		if resource == nil {
			return nil
		}
		return raml.Add(method, route, resource)
	})
}

// Make this configurable.
func githubURL(info chi.FuncInfo) string {
	str := fmt.Sprintf("https://%v#L%v", info.File, info.Line)
	return strings.Replace(str, "github.com/pressly/api/", "github.com/pressly/api/blob/master/", 1)
}

func DeveloperDocs(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) (*Resource, error) {
	if method == "*" {
		return nil, nil
	}

	info := chi.GetFuncInfo(handler)
	pkg := info.Pkg[strings.LastIndex(info.Pkg, "/")+1:]

	desc := ""
	parts := strings.SplitAfterN(info.Comment, ".", -1)
	if len(parts) > 0 {
		desc += fmt.Sprintf("<h3>%s</h3>\n", parts[0])
	}
	if len(parts) > 1 {
		desc += fmt.Sprintf("%s\n", parts[1])
	}

	desc += fmt.Sprintf("\n\n---\n\n⇩ HTTP Request<br />\n")

	if len(middlewares) > 0 {
		func() {
			for i, mw := range middlewares {
				mwInfo := chi.GetFuncInfo(mw)
				mwPkg := mwInfo.Pkg[strings.LastIndex(mwInfo.Pkg, "/")+1:]
				desc += fmt.Sprintf("%v↳ [%v.**%v**](%v)<br />\n", strings.Repeat("&nbsp;", 2*(i+1)), mwPkg, mwInfo.Func, githubURL(mwInfo))
				if i == len(middlewares)-1 {
					desc += fmt.Sprintf("%v↳<br />\n", strings.Repeat("&nbsp;", 2*(i+2)))
					desc += fmt.Sprintf("%v[%v.**%v**](%v)<br />\n", strings.Repeat("&nbsp;", 2*(i+3)), pkg, info.Func, githubURL(info))
					desc += fmt.Sprintf("%v↵<br />\n", strings.Repeat("&nbsp;", 2*(i+2)))
				}
				defer func(i int) {
					desc += fmt.Sprintf("%v↵ [%v.**%v**](%v)<br />\n", strings.Repeat("&nbsp;", 2*(i+1)), mwPkg, mwInfo.Func, githubURL(mwInfo))
				}(i)
			}
		}()
	} else {
		desc += fmt.Sprintf("TODO")
	}

	resource := &Resource{
		Description: desc,
		Responses:   Responses{},
	}

	switch method {
	case "POST":
		resource.Responses[201] = Response{}
	case "GET", "PUT":
		resource.Responses[200] = Response{}
	case "DELETE":
		resource.Responses[204] = Response{}
	}

	return resource, nil
}

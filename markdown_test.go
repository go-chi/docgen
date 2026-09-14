package docgen

import "testing"

func TestGithubSourceURLStripsModuleVersion(t *testing.T) {
	md := &MarkdownDoc{
		Opts: MarkdownOpts{
			ProjectPath: "github.com/example/app",
			URLMap: map[string]string{
				"github.com/go-chi/chi/v5": "https://github.com/go-chi/chi/blob/master/",
			},
		},
	}

	file := "github.com/go-chi/chi/v5@v5.0.11/middleware/logger.go"
	got := md.githubSourceURL(file, 39)
	want := "https://github.com/go-chi/chi/blob/master/middleware/logger.go#L39"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestGithubSourceURLStripsModuleVersionFromModCachePath(t *testing.T) {
	md := &MarkdownDoc{
		Opts: MarkdownOpts{
			ProjectPath: "github.com/example/app",
			URLMap: map[string]string{
				"github.com/go-chi/chi/v5": "https://github.com/go-chi/chi/blob/master/",
			},
		},
	}

	file := "/Users/me/go/pkg/mod/github.com/go-chi/chi/v5@v5.0.11/middleware/logger.go"
	got := md.githubSourceURL(file, 39)
	want := "https://github.com/go-chi/chi/blob/master/middleware/logger.go#L39"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

package docgen

import "testing"

func TestGithubSourceURLStripsModuleVersion(t *testing.T) {
	md := &MarkdownDoc{
		Opts: MarkdownOpts{
			ProjectPath: "github.com/go-chi/chi/v5",
			URLMap: map[string]string{
				"github.com/go-chi/chi/v5": "https://github.com/go-chi/chi/tree/master",
			},
		},
	}

	file := "github.com/go-chi/chi/v5@v5.0.11/middleware/logger.go"
	got := md.githubSourceURL(file, 39)
	want := "https://github.com/go-chi/chi/tree/master/middleware/logger.go#L39"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

package gvtunused

import (
	"os/exec"
	"testing"
)

func TestGoListImports(t *testing.T) {
	testCases := []struct {
		desc    string
		output  string
		imports []string
	}{
		{
			desc:    "single import list",
			output:  "github.com/slomek/gvt-unused",
			imports: []string{"github.com/slomek/gvt-unused"},
		}, {
			desc: "multiple import list",
			output: `github.com/slomek/gvt-unused/a
			github.com/slomek/gvt-unused/b
			github.com/slomek/gvt-unused/c
			`,
			imports: []string{
				"github.com/slomek/gvt-unused/a",
				"github.com/slomek/gvt-unused/b",
				"github.com/slomek/gvt-unused/c"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			cmd := exec.Command("echo", tC.output)
			imps, err := GoListImports(cmd)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if len(imps) != len(tC.imports) {
				t.Fatalf("Expected %d imports, got: %d", len(tC.imports), len(imps))
			}

			for i, imp := range imps {
				if tC.imports[i] != imp {
					t.Errorf("Expected import #%d to be %s, got: %s", i, tC.imports[i], imp)
				}
			}
		})
	}
}

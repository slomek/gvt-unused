package gvtunused

import "testing"

func TestListManifestDeps(t *testing.T) {
	testCases := []struct {
		desc         string
		manifestPath string
		dependencies []string
	}{
		{
			desc:         "parse single-dependency manifest file",
			manifestPath: "testdata/single.manifest",
			dependencies: []string{"github.com/pkg/errors"},
		}, {
			desc:         "parse multiple-dependency manifest file",
			manifestPath: "testdata/multiple.manifest",
			dependencies: []string{
				"github.com/pkg/errors",
				"github.com/gorilla/mux"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			deps, err := ListManifestDeps(tC.manifestPath)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if len(deps) != len(tC.dependencies) {
				t.Fatalf("Expected %d dependencies, got: %d", len(tC.dependencies), len(deps))
			}

			for i, imp := range deps {
				if tC.dependencies[i] != imp {
					t.Errorf("Expected dependency #%d to be %s, got: %s", i, tC.dependencies[i], imp)
				}
			}
		})
	}
}

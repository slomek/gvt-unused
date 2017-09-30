package gvtunused

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

type manifestFile struct {
	Dependencies []struct {
		ImportPath string `json:"importpath,omitempty"`
	} `json:"dependencies,omitempty"`
}

func ListManifestDeps(manifestPath string) ([]string, error) {
	raw, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read manifest file")
	}

	var m manifestFile
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal manifest file")
	}

	deps := make([]string, 0, len(m.Dependencies))
	for _, mI := range m.Dependencies {
		deps = append(deps, mI.ImportPath)
	}

	return deps, nil
}

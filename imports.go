package gvtunused

import (
	"log"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

func UsedDependencies() ([]string, error) {
	log.Println("  ↳ Looking in the sources")
	cmd := exec.Command("go", "list", "-f", "{{ join .Deps \"\\n\" }}", ".")
	deps, err := GoListImports(cmd)
	if err != nil {
		return nil, errors.Wrap(err, "cmd0")
	}

	log.Println("  ↳ Looking in the test sources")
	cmd = exec.Command("go", "list", "-f", "{{ join .TestImports \"\\n\" }}", ".")
	testDeps, err := GoListImports(cmd)
	if err != nil {
		return nil, errors.Wrap(err, "cmd0")
	}
	deps = append(deps, testDeps...)

	log.Println("  ↳ Filtering non-standard imports...")
	var used []string
	for _, d := range deps {
		cmd = exec.Command("go", "list", "-f", "{{if .Standard}}{{else}}{{.ImportPath}}{{end}}", d)
		imps, err := GoListImports(cmd)
		if err != nil {
			return nil, err
		}
		used = append(used, imps...)
	}

	rootPath, err := projectRoot()
	if err != nil {
		return nil, errors.Wrap(err, "failed to find project root path")
	}

	var importPaths []string
	for _, u := range used {
		importPath := strings.Replace(u, rootPath, "", -1)
		importPaths = append(importPaths, importPath)
	}

	return importPaths, nil
}

func projectRoot() (string, error) {
	cmd := exec.Command("go", "list", "./...")
	imps, err := GoListImports(cmd)
	if err != nil {
		return "", errors.Wrap(err, "failed to list imports")
	}
	return imps[0], nil
}

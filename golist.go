package gvtunused

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

func GoListImports(cmd *exec.Cmd) ([]string, error) {
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, errors.Wrap(err, "failed to list dependencies")
	}

	outputString := string(out.Bytes())
	outputLines := strings.Split(outputString, "\n")

	var imports []string
	for _, oL := range outputLines {
		oL = strings.TrimSpace(oL)
		if oL != "" {
			imports = append(imports, oL)
		}
	}

	return imports, nil
}

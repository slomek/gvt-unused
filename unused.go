package gvtunused

import "strings"

func UnusedDependencies(manifestDeps, usedDeps []string) []string {
	var unused []string
	for _, dep := range manifestDeps {
		if !contains(usedDeps, dep) {
			unused = append(unused, dep)
		}
	}
	return unused
}

func contains(deps []string, dep string) bool {
	for _, depItem := range deps {
		if strings.Contains(depItem, dep) {
			return true
		}
	}
	return false
}

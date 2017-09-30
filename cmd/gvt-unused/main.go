package main

import (
	"flag"
	"log"

	"github.com/slomek/gvt-unused"
)

func main() {
	manifest := flag.String("manifest", "./vendor/manifest", "path to gvt manifest file")
	flag.Parse()

	log.SetFlags(0)

	log.Println("‣ Listing dependencies from manifest file...")
	manifestDeps, err := gvtunused.ListManifestDeps(*manifest)
	if err != nil {
		log.Fatalf("Failed to list manifest dependencies: %v", err)
	}

	log.Println("‣ Listing dependencies from source code...")
	usedDeps, err := gvtunused.UsedDependencies()
	if err != nil {
		log.Fatalf("Failed to list used dependencies: %v", err)
	}

	log.Println("‣ Comparing manifest and used dependencies...")
	unusedDeps := gvtunused.UnusedDependencies(manifestDeps, usedDeps)

	if len(unusedDeps) == 0 {
		log.Println("✔ There are no unused gvt dependencies.")
	} else {
		log.Println()
		log.Printf("✘ Found %d unused dependencies, to remove them execute:\n", len(unusedDeps))
		for _, ud := range unusedDeps {
			log.Printf("gvt delete %s", ud)
		}
		log.Println()
	}
}

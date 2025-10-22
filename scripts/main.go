package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/The-DevOps-Daily/terraform-provider-validatefx/internal/functions"
)

func main() {
	generateDocs := exec.Command("tfplugindocs", "generate")
	generateDocs.Stdout = os.Stdout
	generateDocs.Stderr = os.Stderr
	generateDocs.Env = os.Environ()
	generateDocs.Dir = filepath.Clean(".")

	if err := generateDocs.Run(); err != nil {
		log.Fatal(err)
	}

	if err := functions.UpdateReadmeFunctionsTable(context.Background(), "README.md"); err != nil {
		log.Fatal(err)
	}
}

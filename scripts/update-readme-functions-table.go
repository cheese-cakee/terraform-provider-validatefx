package main

import (
	"context"
	"log"

	"github.com/The-DevOps-Daily/terraform-provider-validatefx/internal/functions"
)

func main() {
	if err := functions.UpdateReadmeFunctionsTable(context.Background(), "README.md"); err != nil {
		log.Fatal(err)
	}
}

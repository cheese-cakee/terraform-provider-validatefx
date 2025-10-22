package functions

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// UpdateReadmeFunctionsTable rewrites the Available Functions table in README.md based on current function docs.
func UpdateReadmeFunctionsTable(ctx context.Context, readmePath string) error {
	docs, err := AvailableFunctionDocs(ctx)
	if err != nil {
		return err
	}

	var tableBuilder strings.Builder
	tableBuilder.WriteString("| Function | Description |\n")
	tableBuilder.WriteString("| -------------------------- | ------------------------------------------------ |\n")
	for _, doc := range docs {
		summary := doc.Summary
		if summary == "" {
			summary = doc.Description
		}
		tableBuilder.WriteString(fmt.Sprintf("| `%s` | %s |\n", doc.Name, summary))
	}

	contents, err := os.ReadFile(readmePath)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`(?s)(## 🧩 Available Functions\n\n)(\|.*?\|\n\|.*?\|\n)(.*?)(\n\n---)`)

	replacement := []byte(fmt.Sprintf("$1%s$4", tableBuilder.String()))

	updated := re.ReplaceAll(contents, replacement)

	if bytes.Equal(contents, updated) {
		return nil
	}

	return os.WriteFile(readmePath, updated, 0o644)
}

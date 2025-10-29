#!/usr/bin/env bash
set -euo pipefail

REPO="The-DevOps-Daily/terraform-provider-validatefx"
ISSUES_FILE="issues.json"
DEFAULT_LABEL="hacktoberfest"

if ! command -v gh >/dev/null 2>&1; then
  echo "gh CLI is required but not found in PATH." >&2
  exit 1
fi

if [ ! -f "$ISSUES_FILE" ]; then
  echo "Issues file $ISSUES_FILE not found." >&2
  exit 1
fi

echo "Ensuring default label '$DEFAULT_LABEL' exists..."
if ! gh label list --repo "$REPO" --limit 1000 | cut -f1 | grep -Fxq "$DEFAULT_LABEL"; then
  gh label create "$DEFAULT_LABEL" --repo "$REPO" --description "Hacktoberfest participation" --color D22AE4
fi

tmp=$(mktemp)
cat "$ISSUES_FILE" | jq -c '.[]' > "$tmp"

while IFS= read -r issue; do
  title=$(echo "$issue" | jq -r '.title')
  body=$(echo "$issue" | jq -r '.body')
  labels=$(echo "$issue" | jq -r '.labels | @csv' | tr -d '"')

  IFS=',' read -ra label_array <<< "$labels"
  label_args=()
  for label in "${label_array[@]}"; do
    label=${label//,/}
    label=${label//\"/}
    label_trimmed=$(echo "$label" | xargs)
    if [ -n "$label_trimmed" ]; then
      if ! gh label list --repo "$REPO" --limit 1000 | cut -f1 | grep -Fxq "$label_trimmed"; then
        echo "Label '$label_trimmed' does not exist; creating it."
        gh label create "$label_trimmed" --repo "$REPO" --color AAAAAA
      fi
      label_args+=("--label" "$label_trimmed")
    fi
  done

  label_args+=("--label" "$DEFAULT_LABEL")

  echo "Creating issue: $title"
  gh issue create --repo "$REPO" --title "$title" --body "$body" "${label_args[@]}"
done < "$tmp"

rm -f "$tmp"

#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
CHANGELOG="$REPO_ROOT/CHANGELOG.md"

if ! command -v git >/dev/null 2>&1; then
  echo "git command not found" >&2
  exit 1
fi

latest_tag=$(git describe --tags --abbrev=0 2>/dev/null || echo "")
if [[ -z "$latest_tag" ]]; then
  echo "No tags found. Aborting changelog update." >&2
  exit 1
fi

commits=$(git log "${latest_tag}..HEAD" --merges --pretty=format:"%H")
if [[ -z "$commits" ]]; then
  echo "No commits found since $latest_tag."
  exit 0
fi

regular_changes=()
dependency_updates=()

while IFS= read -r merge_commit; do
  title=$(git log -1 --pretty=format:"%s" "$merge_commit")
  body=$(git log -1 --pretty=format:"%b" "$merge_commit")

  issue=""
  if [[ "$body" =~ ([#][0-9]+) ]]; then
    issue=" ${BASH_REMATCH[1]}"
  fi

  short_sha=${merge_commit:0:7}
  entry="- ${title} (${short_sha})${issue}"

  if [[ "${title,,}" == *"dependabot"* ]]; then
    dependency_updates+=("$entry")
  else
    regular_changes+=("$entry")
  fi

done <<< "$commits"

new_section="## [Unreleased]\n\n"

if ((${#regular_changes[@]} > 0)); then
  new_section+="### Changes\n\n"
  for item in "${regular_changes[@]}"; do
    new_section+="${item}\n"
  done
  new_section+="\n"
fi

if ((${#dependency_updates[@]} > 0)); then
  new_section+="### Dependency Updates\n\n"
  for item in "${dependency_updates[@]}"; do
    new_section+="${item}\n"
  done
  new_section+="\n"
fi

if [[ -f "$CHANGELOG" ]]; then
  tmp=$(mktemp)
  {
    read -r first_line < "$CHANGELOG"
    echo "$first_line"
    echo
    echo -e "$new_section"
    tail -n +2 "$CHANGELOG"
  } > "$tmp"
  mv "$tmp" "$CHANGELOG"
else
  echo -e "# Changelog\n\n$new_section" > "$CHANGELOG"
fi

echo "Changelog updated with commits since $latest_tag."

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

section_date=$(date +"%Y-%m-%d")
section_header="## [${section_date}]"

mapfile -t entries < <(git log "${latest_tag}..HEAD" --pretty=format:"%s (%h)" --no-merges)

notes=()
while IFS= read -r merge_commit; do
  title=$(git log -1 --pretty=format:"%s" "$merge_commit")
  body=$(git log -1 --pretty=format:"%b" "$merge_commit")

  issue=""
  if [[ "$body" =~ ([#][0-9]+) ]]; then
    issue=" ${BASH_REMATCH[1]}"
  fi

  short_sha=${merge_commit:0:7}
  notes+=("- ${title} (${short_sha})${issue}")

  mapfile -t non_merge < <(git log "$merge_commit^1..$merge_commit^2" --pretty=format:"%s (%h)" --no-merges)
  for n in "${non_merge[@]}"; do
    notes+=("  * ${n}")
  done
  notes+=("")
done <<< "$commits"

new_section="$section_header

"
for note in "${notes[@]}"; do
  new_section+="$note
"
done
new_section+="
"

if [[ -f "$CHANGELOG" ]]; then
  tmp=$(mktemp)
  {
    echo "$new_section"
    echo
    cat "$CHANGELOG"
  } > "$tmp"
  mv "$tmp" "$CHANGELOG"
else
  echo -e "$new_section" > "$CHANGELOG"
fi

echo "Changelog updated with commits since $latest_tag."

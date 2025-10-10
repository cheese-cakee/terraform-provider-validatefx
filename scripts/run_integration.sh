#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

cd "$REPO_ROOT"

docker compose build terraform
# Ensure terraform state is cleaned after each run
trap 'docker compose -f docker-compose.yml down -v' EXIT

docker compose up --abort-on-container-exit

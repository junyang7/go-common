#!/usr/bin/env bash

set -e

FILE="./VERSION"
TIMESTAMP="$(date +%Y%m%d%H%M%S)"
VERSION="$(cat "$FILE")"
IFS='.' read -r MAJOR MINOR PATCH <<< "$VERSION"
NEW_VERSION="${MAJOR}.${MINOR}.${TIMESTAMP}"
echo "$NEW_VERSION" > "$FILE"

export https_proxy=http://127.0.0.1:7897 http_proxy=http://127.0.0.1:7897 all_proxy=socks5://127.0.0.1:7897
git add .
git commit -m "$NEW_VERSION"
git push

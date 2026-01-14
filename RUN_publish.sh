#!/usr/bin/env bash

set -e

PATH="./VERSION"
TIMESTAMP="$(/bin/date +%Y%m%d%H%M%S)"
VERSION="$(/bin/cat "$PATH")"
IFS='.' read -r MAJOR MINOR PATCH <<< "$VERSION"
NEW_VERSION="${MAJOR}.${MINOR}.${TIMESTAMP}"
echo "$NEW_VERSION" > "$PATH"

export https_proxy=http://127.0.0.1:7897 http_proxy=http://127.0.0.1:7897 all_proxy=socks5://127.0.0.1:7897
/usr/bin/git add .
/usr/bin/git commit -m "$NEW_VERSION"
/usr/bin/git push

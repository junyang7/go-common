#!/usr/bin/env bash

set -e

VERSION="2.0.0"
PATH="./VERSION"

new_version="${VERSION}+$(/bin/date +%Y%m%d%H%M%S)"
echo "$new_version" > "$PATH"

export https_proxy=http://127.0.0.1:7897 http_proxy=http://127.0.0.1:7897 all_proxy=socks5://127.0.0.1:7897
/usr/bin/git add . && /usr/bin/git commit -m "$new_version" && /usr/bin/git push

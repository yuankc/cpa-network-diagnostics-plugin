#!/usr/bin/env bash
set -euo pipefail

goos="${GOOS:-$(go env GOOS)}"
goarch="${GOARCH:-$(go env GOARCH)}"
ext=".so"
if [[ "$goos" == "darwin" ]]; then
  ext=".dylib"
elif [[ "$goos" == "windows" ]]; then
  ext=".dll"
fi

out_dir="dist/$goos/$goarch"
out_file="$out_dir/diagnostics$ext"
mkdir -p "$out_dir"
CGO_ENABLED=1 go build -buildmode=c-shared -trimpath -ldflags "-s -w" -o "$out_file" .
echo "Built $out_file"


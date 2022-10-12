#!/bin/bash
find . -maxdepth 1 -type f -name "*.zip" -exec ./convert_to_zstd.sh {} \;
find . -maxdepth 1 -type f -name "godot-docs-html-*.tar.zst" -exec ./build.sh {} \;

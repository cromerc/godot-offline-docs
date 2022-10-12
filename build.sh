#!/bin/bash
shopt -s dotglob

filename="${1}"
newfilename="${filename%.*.*}"
version="${newfilename:18}"

echo "${version}" > version.txt

mkdir godot-docs-html
tar -xpmf "${filename}" -C godot-docs-html

pushd godot-docs-html
find . -type f -exec gzip {} \;
popd

go build -ldflags="-s -w"
upx godot-docs
mv godot-docs godot-docs-"${version}"
tar --zstd -cpf godot-docs-"${version}".tar.zst godot-docs-"${version}"
rm -f godot-docs-"${version}"

rm -rf godot-docs-html

rm -f version.txt

mv "${filename}" html
mv godot-docs-"${version}".tar.zst docs

shopt -u dotglob

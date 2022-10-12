#!/bin/bash
shopt -s dotglob

filename="${1}"
newfilename="${filename%.*}"

echo "Converting ${filename} to ${newfilename}.tar.zstd"
unzip -q "${filename}" -d "${newfilename}"

pushd "${newfilename}"
tar --zstd -cpf ../"${newfilename}.tar.zst" *
popd

rm -rf "${newfilename}"
rm -f "${filename}"

shopt -u dotglob

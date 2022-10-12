# Godot offline docs

The Godot offline docs are an all-in-one binary that when run makes the documentation available locally in your browser. There are several benefits to this program. The documents being offline load instantly while moving page to page in the docs and can be used when you don't have a stable internet connect. The size is small because all documentation is gzipped inside the binary. It is portable and can be run on pretty much any distro that has GTK3 installed.

## Requirements

### Compile
- Golang 1.18
- A copy of the html godot documentation en .zip format.
- upx
- GTK3

### Runtime
- GTK3

### Binaries
Linux binaries are available to download on the release page containing the offline documentation.

Right now I don't have working windows binaries of this, but if anyone wants to help get it working and maybe even automated, please open a pull request.
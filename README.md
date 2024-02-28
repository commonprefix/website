# commonprefix website

## Instructions

1. [Install the go programming language](https://go.dev/doc/install)
2. Run `go run .` (or `go build . && ./website`)
3. `./public` will have the static artifacts of the website

> Run `go run . --serve` or `./website --serve` for building and serving the files. Rerun every time source files change.

## Static files

`/public/static` holds all static files: stylesheets, images, fonts and documents. For team member pictures, add a smaller version with a suffix of `_w150` (e.g. `john_doe_w150.jpg`). You can run `sips -Z 150 john_doe_w150.jpg` to downsize the image.

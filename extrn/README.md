# extrn

Extension renaming utility.

## Usage

```
Usage of ./extrn:
  -extIn="mov": Input file extension
  -extOut="mpg": Output file extension
```

If no additional arguments are given, the current working directory will be
processed, otherwise all listed directories will be scanned and all
files will be processed.

## Dependencies

 * [Go](http://golang.org) - for compilation

## Building

```
go build
```

To cross-compile, make sure that you have installed and setup
``golang-crosscompile`` first. Then run

```
make
```


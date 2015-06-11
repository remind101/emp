# emp

A CLI for Empire.

## Installation

```console
$ go get -u github.com/remind101/emp
```

## Usage

The basic usage of emp is:

```
Usage: emp <command> [-a app] [options] [arguments]
```

## Development

emp requires Go 1.2 or later and uses [Godep](https://github.com/kr/godep) to manage dependencies.

	$ cd emp
	$ vim main.go
	$ godep go build
	$ ./emp apps

Please follow the [contribution guidelines](./CONTRIBUTING.md) before submitting
a pull request.

[go-install]: http://golang.org/doc/install "Golang installation"

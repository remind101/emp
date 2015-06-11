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

## Config

emp allows some configuration using git config.

### Use strict flag ordering / disable interspersed flags and non-flag arguments

Enable:

```
$ git config --global --bool emp.strict-flag-ordering true
```

Disable:

```
$ git config --global --unset emp.strict-flag-ordering
```

## Plugins

emp currently has a minimal plugin system. It may see substantial changes in the future, and those changes may break existing plugins or change the architecture at any time. Use this functionality at your own risk.

Plugins are executables located in HKPATH or, if HKPATH does not exist, in /usr/local/lib/emp/plugin. They are executed when emp does not know command X and an installed plugin X exists. The special case default plugin will be executed if emp has no command or installed plugin named X.

emp will set these environment variables for a plugin:

* EMPIRE_API_URL - The url containing the username, password, and host to the api endpoint.
* HKAPP - The app as determined by the git heroku remote, if available.
* HKUSER - The username from either EMPIRE_API_URL or .netrc
* HKPASS - The password from either EMPIRE_API_URL or .netrc
* HKHOST - The hostname for the API endpoint

## Development

emp requires Go 1.2 or later and uses [Godep](https://github.com/kr/godep) to manage dependencies.

	$ cd emp
	$ vim main.go
	$ godep go build
	$ ./emp apps

Please follow the [contribution guidelines](./CONTRIBUTING.md) before submitting
a pull request.

[go-install]: http://golang.org/doc/install "Golang installation"

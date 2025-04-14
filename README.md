# Golang application version

This package exports `func Version() (version string)` which returns a version number for the application in which it is included.
If possible, the version number returned will follow the
[golang semantic versioning model](https://go.dev/doc/modules/version-numbers). The version can be overridden by
injecting the version when building the application.

Since golang version 1.24, version handling changed and because of that golang version 1.24 is the minimum golang
version that it is possible to use this package with.

This package does not import any other packages.

## Installation

```shell
go get github.com/utvecklat/version
```

## Usage

Call `version.Version()` and use the returned string as the application version number.

See the example at [example/example.go](example/example.go)

There is also an example application at [github.com/utvecklat/versionexample](github.com/utvecklat/versionexample) with 
some more detailed examples of how to set version numbers and build applications. That project can also demonstrate 
how to install or run an application with this package included.

### Exports

`func Version() (version string)`

### Version types 

* (devel)
    * Code is not in a VCS tree
    * Code is executed with `go run`
* v1.2.3
    * The source code is in a VCS and current commit is tagged with `v1.2.3`
    * There are no changes since the last commit to the VCS
    * Code compiled with `go install` or `go build`
* v1.2.3+dirty
    * Code is in a VCS and latest commit is tagged with `v1.2.3` but there are changes that is not commited
    * Code compiled with `go build`
* v1.2.4-0.20250408082215-7188c2627f0a
    * The code is in a VCS and the latest tag is `v1.2.3` but there are commits since that tag.
    * There are no changes since last commit to the VCS
    * Code compiled with `go install` or `go build`
* v1.2.4-0.20250408082215-7188c2627f0a+dirty
    * The code is in a VCS and the latest tag is `v1.2.3` but there are commits since that tag.
    * There are changes that have not been commited to the VCS
    * Code compiled with `go build`

### Version override

To override the version number, inject a version when building the code. If this is done no other version number scheme
will be used.

```shell
go build -trimpath -ldflags "-X github.com/utvecklat/version.versionOverride=v1.2.3"
```

This is primarily for projects that do not use version control systems with version tags or alternative build methods.

## License and Code of Conduct

This package is released under the [MIT License](LICENSE).
Please follow the [code of conduct](CODE_OF_CONDUCT.md) when you interact with the project contributors.

## References

### Golang version numbers

Read more about golang versions here:

https://go.dev/doc/modules/version-numbers

### Changes in golang 1.24

main.Version changed in version 1.24

[Go 1.24 Release Notes](https://tip.golang.org/doc/go1.24)

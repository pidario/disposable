# Disposable email domains helpers
[![GoDoc](https://godoc.org/github.com/pidario/disposable?status.svg)](https://godoc.org/github.com/pidario/disposable) [![Build Status](https://travis-ci.com/pidario/disposable.svg?branch=master)](https://travis-ci.com/pidario/disposable) [![Go Report Card](https://goreportcard.com/badge/github.com/pidario/disposable)](https://goreportcard.com/report/github.com/pidario/disposable)

Disposable email domains helpers (based on [ivolo/disposable-email-domains](https://github.com/ivolo/disposable-email-domains))

Due to Go naming conventions (hyphens in package name should be avoided) I preferred to create this repository instead of submitting a pull request but the idea is to keep `list/index.json` up-to-date with parent repository.
I did not find `wildcard.json` file useful so I did not include it but maybe in the future I will.

As of now, only Go helper is present but I mean to add helpers for other languages and publish this repository to the main package managers (such as npm).
# Go
## Development
### Change version
it should be the latest version of the parent library (ivolo/disposable-email-domains);
important: don't prepend 'v'
```
echo "1.0.56" > version
```
### Generating asset file
necessary each time `list/index.json` is updated
```
make generate
```
### Testing
```
make test
```
## Usage
```go
import (
	"fmt"
	
	"github.com/pidario/disposable"
)

func main() {
	domainChecker := disposable.NewDomainChecker()
	if domainChecker.Error != nil {
		// in this unlikely scenario (it means the file list/index.json
		// is not present or is not valid JSON when you install the package)
		// IsDisposable always returns false
	}
	isBL := domainChecker.IsDisposable("mailinator.com")
	fmt.Println(isBL) // true
	// you can also inspect the original list
	// and create your custom function to check disposability
	fmt.Println(domainChecker.List)
}
```
# Contributing
Of course PR are more than welcome. But as I mentioned, since the list itself depends on another repository, I will consider only those that add or modify helper codes and/or tests (in any language).
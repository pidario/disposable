# Disposable email domains helpers
Disposable email domains helpers (based on [ivolo/disposable-email-domains](https://github.com/ivolo/disposable-email-domains))
Due to Go naming conventions (hyphens in package name should be avoided) I preferred to create this repository instead of sumitting a pull request but the idea is to keep `index.json` up-to-date with parent repository.
I did not find useful `wildcard.json` file so it is not included but maybe in the future I will.
As of now, only Golang helper is present but I mean to add helpers for other languages and publish this repository to the main package managers (such as npm).
# Go
## Installation
```
go get -u -v github.com/pidario/disposable
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
		// in this unlikely scenario (it means the file index.json
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
Of course PR are more than welcome. But as I mentioned, the list itself depends on another repository so I will consider only those that modify the helper code and/or tests (Go or other languages if any). Of course unit tests should be added too.
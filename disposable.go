// Package disposable contains a helper function used to check whether an email domain
// is to be considered disposable
package disposable

import (
	"bufio"
	"encoding/json"
	"sort"
	"strings"
)

//go:generate go run vfsdata_generate.go

// Domains contains the list of the disposable domains and error (if any)
type Domains struct {
	List  []string
	Error error
}

func readList() ([]string, error) {
	var list []string
	file, err := asset.Open("index.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	err = json.NewDecoder(rd).Decode(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// NewDomainChecker initializes a Domains object
func NewDomainChecker() (d Domains) {
	list, err := readList()
	if err != nil {
		d.Error = err
		return
	}
	d.List = list
	return
}

// IsDisposable checks if the provided domain is contained in list
func (d *Domains) IsDisposable(domain string) bool {
	// d.Error != nil means that some sort of reading error happened, so domain should NOT be considered disposable
	if d.Error != nil {
		return false
	}
	domain = strings.ToLower(domain)
	index := sort.SearchStrings(d.List, domain)
	if index < len(d.List) && strings.ToLower(d.List[index]) == domain {
		return true
	}
	return false
}

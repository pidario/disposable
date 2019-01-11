// Package disposable a helper function to check whether an email domain
// is to be considered disposable
package disposable

import (
	"bufio"
	"encoding/json"
	"sort"
	"strings"
)

//go:generate go run vfsdata_generate.go

// Domains contains the list of the blacklisted domains and eventual error
type Domains struct {
	List  []string
	Error error
}

func read() ([]string, error) {
	var content []string
	var b []byte
	file, err := asset.Open("index.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Bytes()
		for _, _t := range t {
			b = append(b, _t)
		}
	}
	err = json.Unmarshal(b, &content)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// NewDomainChecker initializes a Domains object
func NewDomainChecker() (d Domains) {
	content, err := read()
	if err != nil {
		d.Error = err
		return
	}
	d.List = content
	return
}

// IsDisposable checks if the provided domain is contained in blacklist
func (d *Domains) IsDisposable(domain string) bool {
	// d.Error != nil means that some sort of reading error happended, so domain should NOT be considered blacklisted
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

package disposable

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/gobuffalo/packr"
)

// Domains contains the list of the blacklisted domains and eventual error
type Domains struct {
	List  []string
	Error error
}

func read() ([]string, error) {
	box := packr.NewBox(".")
	var content []string
	file, err := box.Find("index.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &content)
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

package disposable

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func newDisposableDomains(t *testing.T) Domains {
	D := NewDomainChecker()
	if D.Error != nil {
		t.Fatal(D.Error)
	}
	if len(D.List) == 0 {
		t.Fatal("file index.json is empty")
	}
	return D
}

func testDisposability(domains []string, expected bool, t *testing.T) {
	D := newDisposableDomains(t)
	not := ""
	if !expected {
		not = " NOT"
	}
	for _, domain := range domains {
		if res := D.IsDisposable(domain); res != expected {
			t.Errorf("domain %s%s expected to be disposable but found %v", domain, not, res)
		}
	}
}

func TestDisposable(t *testing.T) {
	domains := []string{
		"0-mail.com",
		"1CE.US",
		"mailiNator.com",
		"zzz.COM",
	}
	testDisposability(domains, true, t)
}

func TestNotDisposable(t *testing.T) {
	domains := []string{
		"gmail.com",
		"yahoo.com",
	}
	testDisposability(domains, false, t)
}

func TestDuplicates(t *testing.T) {
	D := newDisposableDomains(t)
	m := map[string]bool{}
	for _, domain := range D.List {
		if m[domain] {
			t.Errorf("%s is duplicate", domain)
		} else {
			m[domain] = true
		}
	}
}

func TestIsSorted(t *testing.T) {
	D := newDisposableDomains(t)
	var sortedSlice []string
	for _, v := range D.List {
		sortedSlice = append(sortedSlice, v)
	}
	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i] < sortedSlice[j]
	})
	if len(sortedSlice) != len(D.List) {
		t.Fatal("an error occurred while checking whether the list is sorted")
	}
	for i, v := range D.List {
		if v != sortedSlice[i] {
			t.Fatalf("list is not sorted: index %d, value %s", i, v)
		}
	}
}

func TestIsLowerCase(t *testing.T) {
	D := newDisposableDomains(t)
	var lowerCaseSlice []string
	for _, v := range D.List {
		lowerCaseSlice = append(lowerCaseSlice, strings.ToLower(v))
	}
	if len(lowerCaseSlice) != len(D.List) {
		t.Fatal("an error occurred while checking whether the list is lowercase")
	}
	for i, v := range D.List {
		if v != lowerCaseSlice[i] {
			t.Errorf("list is not lowercase: index %d, value %s", i, v)
		}
	}
}

func TestError(t *testing.T) {
	D := NewDomainChecker()
	D.Error = fmt.Errorf("manual error")
	if check := D.IsDisposable("mailinatar.com"); check {
		t.Fatal("mailinatar.com is disposable but it should be flagged as non-disposable because the error is set")
	}
}

package scan

import (
	"errors"
	"fmt"
	"sort"
)

var (
	ErrExists    = errors.New("host already in list.")
	ErrNotExists = errors.New("Host not in list.")
)

// slice of hosts to be scanned
type HostsList struct {
	Hosts []string
}

// search for a host in slice
func (hl *HostsList) Search(host string) (bool, int) {
	sort.Strings(hl.Hosts)
	i := sort.SearchStrings(hl.Hosts, host)

	// check if 'host' is in slice after sorting
	if i < len(hl.Hosts) && hl.Hosts[i] == host {
		return true, i
	}
	return false, -1 // -1 to indicate not in slice
}

func (hl *HostsList) Add(host string) error {
	// search for host (dont want to add x2)
	if found, _ := hl.Search(host); found {
		return fmt.Errorf("%w: %s", ErrExists, host)
	}
	// otherwise append to list
	hl.Hosts = append(hl.Hosts, host)
	return nil
}

func (hl *HostsList) Remove(host string) error {
	found, index := hl.Search(host)
	if !found {
		return fmt.Errorf("%w: %s", ErrExists, host)
	}
	hl.Hosts = append(hl.Hosts[:index], hl.Hosts[index+1:]...)
	return nil
}

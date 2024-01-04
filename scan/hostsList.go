package scan

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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

// remove a host from slice
func (hl *HostsList) Remove(host string) error {
	found, index := hl.Search(host)
	if !found {
		return fmt.Errorf("%w: %s", ErrNotExists, host)
	}
	hl.Hosts = append(hl.Hosts[:index], hl.Hosts[index+1:]...)
	return nil
}

// Obtain 'hosts' from a hosts file
func (hl *HostsList) Load(hostFile string) error {
	f, err := os.Open(hostFile)

	// cant open file
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		hl.Hosts = append(hl.Hosts, scanner.Text())
	}
	return nil
}

// Save 'hosts' to a hosts file
func (hl *HostsList) Save(hostFile string) error {
	output := ""
	for _, h := range hl.Hosts {
		output += fmt.Sprintln(h)
	}
	return os.WriteFile(hostFile, []byte(output), 0644) // rw: owner && r: others
}

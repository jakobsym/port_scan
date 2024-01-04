package scan_test

import (
	"errors"
	"pscan/scan"
	"testing"
)

func TestAdd(t *testing.T) {
	// Create test cases
	var testCases = []struct {
		name      string
		host      string
		expectLen int
		expectErr error
	}{
		{"AddNew", "host2", 2, nil},                 // case 1
		{"AddExisting", "host1", 1, scan.ErrExists}, // case 2
	}
	// iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hl := &scan.HostsList{} // create host list

			// Add "host1" to list
			if err := hl.Add("host1"); err != nil {
				t.Fatal(err)
			}
			// Add "host2" to list
			err := hl.Add(tc.host)

			// if error occured; expecting tc.expectErr == nil
			if tc.expectErr != nil {
				// err shouldnt == nil; if tc.expectErr != nil
				if err == nil {
					t.Fatalf("Expected error, got 'nil' instead.\n")
				}
				// if err doesnt match tc.expectErr
				if !errors.Is(err, tc.expectErr) {
					t.Errorf("Expected error: %q, got %q instead.\n", tc.expectErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("Expected nil, got %q instead.\n", err)
			}
			if len(hl.Hosts) != tc.expectLen {
				t.Fatalf("Expected: %q, got %q instead.\n", tc.expectLen, len(hl.Hosts))
			}
			if hl.Hosts[1] != tc.host {
				t.Fatalf("Expected host name: %q, got %q instead.\n", tc.host, hl.Hosts[1])
			}
		})
	}
}

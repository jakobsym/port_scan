package scan_test

import (
	"errors"
	"os"
	"pscan/scan"
	"testing"
)

func TestAdd(t *testing.T) {
	// Create test cases
	testCases := []struct {
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

func TestRemove(t *testing.T) {

	testCases := []struct {
		name      string
		host      string
		expectLen int
		expectErr error
	}{
		{"RemoveExist", "host1", 1, nil},
		{"RemoveNotFound", "host4", 1, scan.ErrNotExists},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hl := &scan.HostsList{}

			// Add hosts to HostList{}
			for _, h := range []string{"host1", "host2"} {
				if err := hl.Add(h); err != nil {
					t.Fatal(err)
				}
			}

			err := hl.Remove(tc.host)

			if tc.expectErr != nil {
				if err == nil {
					t.Fatalf("Expected error, got 'nil' instead.\n")
				}

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
			if hl.Hosts[0] == tc.host {
				t.Fatalf("Host name: %q should not be in the list\n", tc.host)
			}
		})
	}
}

func TestSaveLoad(t *testing.T) {
	hl1 := scan.HostsList{} // Save this hl
	hl2 := scan.HostsList{} // Load saved hl into this hl

	hostName := "host1"
	hl1.Add(hostName)

	tf, err := os.CreateTemp("", "") // uses default dir and file names
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name()) // remove file after operations

	if err := hl1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to the file: %s", err)
	}

	if err := hl2.Load(tf.Name()); err != nil {
		t.Fatalf("Error loading list from file: %s", err)
	}

	if hl1.Hosts[0] != hl2.Hosts[0] {
		t.Errorf("Host %q should match %q host.", hl1.Hosts[0], hl2.Hosts[0])
	}

}

func TestLoadNoFile(t *testing.T) {
	tf, err := os.CreateTemp("", "")

	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	// remove
	if err := os.Remove(tf.Name()); err != nil {
		t.Fatalf("Error removing file: %s", err)
	}
	// create new HostList
	hl := &scan.HostsList{}
	// load the file name (tf.Name())
	if err := hl.Load(tf.Name()); err != nil {
		t.Errorf("Expected no error. Got %q instead\n", err)
	}
}

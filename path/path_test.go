package path

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// TestExists test Exists function
func TestExists(t *testing.T) {
	// Create tmp dir
	tmpDir, err := ioutil.TempDir(os.TempDir(), "util_file_test_")
	if err != nil {
		t.Fatal("Failed to test: failed to create temp dir.")
	}

	// create tmp file
	tmpFile, err := ioutil.TempFile(tmpDir, "test_file_exists_")
	if err != nil {
		t.Fatal("Failed to test: failed to create temp file.")
	}
	tmpFile.Close()

	defer os.RemoveAll(tmpDir)

	tests := []struct {
		name          string
		fileName      string
		expectedExist bool
	}{
		{"file_not_exists", filepath.Join(tmpDir, "file_not_exist"), false},
		{"file_exists", tmpFile.Name(), true},
	}

	for _, test := range tests {
		if realValued, realError := Exists(test.fileName); err != nil {
			t.Fatalf("Failed to test with '%s': %s, '%v'", test.fileName,
				test.name, realError)
		} else if test.expectedExist != realValued {
			t.Fatalf("actual exist was '%v'; "+
				"want '%v'", realValued, test.expectedExist)
		}
	}
}

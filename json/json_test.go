package json

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type jsonData struct {
	Str string `json:"str"`
	Int int    `json:"int"`
}

func TestWriteAndReadJSONForFile(t *testing.T) {
	jd := &jsonData{Str: "string", Int: 8080}
	// when write json to file
	expectedContent := `{"str":"string","int":8080}`

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

	// test WriteJSONToFile
	if err = WriteJSONToFile(tmpFile.Name(), jd); err != nil {
		t.Fatalf("[WriteJSONToFile error]: %s", err)
	}

	contentBytes, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("[ReadFile error]: %s", err)
	}
	// Convert to a string and remove line breaks.
	content := strings.TrimSpace(bytes.NewBuffer(contentBytes).String())

	if content != expectedContent {
		t.Fatalf("actual content was '%v'; "+
			"want '%v'", content, expectedContent)
	}

	// test ReadJSONFromFile
	actualJD := &jsonData{}
	if err = ReadJSONFromFile(tmpFile.Name(), actualJD); err != nil {
		t.Fatalf("[ReadJSONFromFile error]: %s", err)
	}

	if *actualJD != *jd {
		t.Fatalf("actual content was '%v'; "+
			"want '%v'", *actualJD, *jd)
	}
}

package json

import (
	"encoding/json"
	"os"
)

// WriteJSONToFile write v as json to file.
func WriteJSONToFile(path string, v interface{}) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()

	return json.NewEncoder(fp).Encode(v)
}

// ReadJSONFromFile read json file and take into v
func ReadJSONFromFile(path string, v interface{}) error {
	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fp.Close()

	return json.NewDecoder(fp).Decode(v)
}

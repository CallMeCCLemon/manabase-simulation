package reader

import (
	"encoding/json"
	"io"
	"os"
)

func ReadJSONFile[T any](filename string) (T, error) {
	var data T
	file, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer file.Close()

	// Read the file contents
	bytes, err := io.ReadAll(file)
	if err != nil {
		return data, err
	}

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

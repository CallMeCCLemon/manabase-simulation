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

// ReadJSONString Reads a JSON string and returns a struct.
func ReadJSONString[T any](input string) (T, error) {
	// Read the file contents
	var data T

	// Unmarshal the JSON data into the struct
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// WriteJSONString writes an object as JSON to a string.
func WriteJSONString[T any](input T) (string, error) {
	jsonPayload, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	return string(jsonPayload), nil
}

// WriteJSONFile writes an object as JSON to disk.
func WriteJSONFile[T any](filename string, data T) error {
	jsonPayload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonPayload, 0644)
	if err != nil {
		return err
	}
	return nil
}

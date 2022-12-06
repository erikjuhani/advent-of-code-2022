package utils

import (
	"io"
	"os"
)

func ReadInput(path string) (string, error) {
	inputFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer inputFile.Close()

	data, err := io.ReadAll(inputFile)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

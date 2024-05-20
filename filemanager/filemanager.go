package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	
	scanner := bufio.NewScanner(file)
	
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	file.Close()
	err = scanner.Err()
	
	if err != nil {
		return nil, errors.New("failed to read line in file")
	}
	
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	file.Close()

	if err != nil {
		return errors.New("failed to convert data to JSON")
	}

	return nil
}

func New(inputPath, outputPath string) (FileManager) {
	return FileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}
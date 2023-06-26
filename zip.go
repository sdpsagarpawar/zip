package zip

import (
	"archive/zip"
	"bufio"
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ZipOperations interface {
	Zip(ctx context.Context, files map[string][]byte) ([]byte, error)
	Unzip(ctx context.Context, bodyBytes []byte) (map[string][]byte, error)
	ReadFiles(ctx context.Context, dir string) (map[string][]byte, error)
	WriteZip(ctx context.Context, file []byte, zipName string) error
}

type zipOperations struct {
	buffer *bytes.Buffer
}

func NewZipOperations() ZipOperations {
	// buffer needs to be created to store files that are going to zip
	buffer := bytes.Buffer{}
	return &zipOperations{
		buffer: &buffer,
	}
}

// function Zip will take file name as key and file content as byte and return bunch of zip file content as slice of byte
// this can be used to write a file or to upload a file
func (z *zipOperations) Zip(ctx context.Context, files map[string][]byte) ([]byte, error) {
	// Create a buffered writer for the zip file
	zipWriter := zip.NewWriter(bufio.NewWriter(z.buffer))
	defer zipWriter.Close()

	for fileName, fileContent := range files {
		// Create each file with the given name
		zipFile, err := zipWriter.Create(fileName)
		if err != nil {
			return nil, err
		}

		// Write content as zip to the created file
		_, err = zipFile.Write(fileContent)
		if err != nil {
			return nil, err
		}
	}

	// Flush the buffered writer to ensure all data is written
	err := zipWriter.Flush()
	if err != nil {
		return nil, err
	}

	// Return the final bytes as the zip content
	return z.buffer.Bytes(), nil
}

// function Unzip will take the response from HTTP or bytes with the context of a zip file
// and return a map of files with their content and name
// We can write files using this or unmarshal the content of files to required formats such as JSON or XML
func (z *zipOperations) Unzip(ctx context.Context, bodyBytes []byte) (map[string][]byte, error) {
	reader, err := zip.NewReader(bytes.NewReader(bodyBytes), int64(len(bodyBytes)))
	if err != nil {
		return nil, err
	}

	// The reader will return a list of files inside the zip content
	// We can read the content of each file and create a map of files with their content
	response := make(map[string][]byte, len(reader.File))
	for _, file := range reader.File {
		// Open each file to read its content
		fileToRead, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer fileToRead.Close()

		// Read the content of the file
		fileContent, err := ioutil.ReadAll(fileToRead)
		if err != nil {
			return nil, err
		}

		response[file.Name] = fileContent
	}

	return response, nil
}

// function ReadFiles will read files from a directory, make them into a zip, and then return the files' names with their content as a map
// ReadFiles will neglect directories and only read files
func (z *zipOperations) ReadFiles(ctx context.Context, dir string) (map[string][]byte, error) {
	// Read files' details from the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	fileMap := make(map[string][]byte)
	for _, file := range files {
		// Neglect directories and only process files
		if !file.IsDir() {
			filePath := filepath.Join(dir, file.Name())

			// Read the content of the file
			data, err := ioutil.ReadFile(filePath)
			if err != nil {
				return nil, err
			}

			fileMap[file.Name()] = data
		}
	}

	return fileMap, nil
}

// function WriteZip will create a file and write the data inside it
// This can be useful when we want to write data from a file
// We can provide a directory if we want to create the file in a particular directory
func (z *zipOperations) WriteZip(ctx context.Context, file []byte, zipName string) error {
	zipFile, err := os.Create(zipName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(bufio.NewWriter(zipFile))
	defer zipWriter.Close()

	// Create the file in the zip
	zipEntry, err := zipWriter.Create(filepath.Base(zipName))
	if err != nil {
		return err
	}

	// Write the data to the file in the zip
	_, err = zipEntry.Write(file)
	if err != nil {
		return err
	}

	// Flush the buffered writer to ensure all data is written
	err = zipWriter.Flush()
	if err != nil {
		return err
	}

	return nil
}

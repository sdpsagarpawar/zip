//there is no logs in package as we dont want to force users to use logs that we prefer
//user can log errors at their end using any logging library
package zip

import (
	"archive/zip"
	"bytes"
	"context"
	"io/ioutil"
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
	//buffer needs to be created to store files that are going to zip
	buffer := new(bytes.Buffer)
	return &zipOperations{
		buffer: buffer,
	}
}

//function Zip will take file name as key and file content as byte and return bunch of zip file content as slice of byte
//this can be used as to write file or to upload file
func (z *zipOperations) Zip(ctx context.Context, files map[string][]byte) ([]byte, error) {
	//zipWriter needed to write content of buffer to zip
	zipWriter := zip.NewWriter(z.buffer)
	defer zipWriter.Close()
	for fileName, fileContent := range files {
		//create each file with name in input
		zipFile, err := zipWriter.Create(fileName)
		if err != nil {
			return nil, err
		}
		// write content as zip in created file
		_, err = zipFile.Write(fileContent)
		if err != nil {
			return nil, err
		}
	}
	// return final bytes az zip bytes
	return z.buffer.Bytes(), nil
}

//function Unzip will take the response from http or bytes with context of zip file
//and in return it returns map of file with file content and name of file
//we can write file using this or we can unmarshal content of file to required format such as JSON or XML
func (z *zipOperations) Unzip(ctx context.Context, bodyBytes []byte) (map[string][]byte, error) {
	reader, err := zip.NewReader(bytes.NewReader(bodyBytes), int64(len(bodyBytes)))
	if err != nil {
		return nil, err
	}
	//the reader will return us list of files inside zip content
	//we can read content of each file and create map of files with content
	response := make(map[string][]byte, len(reader.File))
	for _, file := range reader.File {
		//for reading content of file we need to open each file
		fileToRead, err := file.Open()
		if err != nil {
			return nil, err
		}
		//the open content can easily readed using ioutil
		fileContent, err := ioutil.ReadAll(fileToRead)
		if err != nil {
			return nil, err
		}
		fileToRead.Close()
		response[file.Name] = fileContent
	}

	return response, nil
}

// function ReadFiles will read files from directory to make it zip and then return files name with content as map
// readfiles will neglect directories and just read files
func (z *zipOperations) ReadFiles(ctx context.Context, dir string) (map[string][]byte, error) {
	// read files detail from directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	fileMap := make(map[string][]byte)
	for _, file := range files {
		// neglect files from directory
		if !file.IsDir() {
			data, err := ioutil.ReadFile(dir + "/" + file.Name())
			if err != nil {
				return nil, err
			}
			//create map of files with file name as directory and file content as bytes
			fileMap[file.Name()] = data
		}
	}
	// return map of file with file name and content
	return fileMap, nil
}

// function WriteZip will create file and write data inside file
// this can be useful when we want write data from file
// we can provide directory if we want to create file in perticuler directory
func (z *zipOperations) WriteZip(ctx context.Context, file []byte, zipName string) error {
	zipWriter := zip.NewWriter(z.buffer)
	defer zipWriter.Close()
	zipFile, err := zipWriter.Create(zipName)
	if err != nil {
		return err
	}
	_, err = zipFile.Write(file)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(zipName, z.buffer.Bytes(), 0777)
}

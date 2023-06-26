# Go Zip Utility

This Go package provides a simple utility for zipping and unzipping files. It allows you to compress multiple files into a single zip archive and extract the contents of a zip archive.

## Features

- Zip files: Compress multiple files into a single zip archive.
- Unzip files: Extract the contents of a zip archive.
- Read files from a directory and create a zip archive.
- Write data to a zip file.

## Installation

Install the package using the following Go command:

```shell
go get github.com/sdpsagarpawar/zip
```
## Uses

```sh
package main

import (
	"context"
	"fmt"

	"github.com/sdpsagarpawar/zip"
)

func main() {
	// Create a new instance of the ZipOperations interface
	zipOps := zip.NewZipOperations()

	// Zip files
	files := map[string][]byte{
		"file1.txt": []byte("This is file 1 content"),
		"file2.txt": []byte("This is file 2 content"),
	}
	zipBytes, err := zipOps.Zip(context.Background(), files)
	if err != nil {
		fmt.Println("Failed to zip files:", err)
		return
	}

	// Unzip files
	unzippedFiles, err := zipOps.Unzip(context.Background(), zipBytes)
	if err != nil {
		fmt.Println("Failed to unzip files:", err)
		return
	}

	// Read files from a directory and create a zip archive
	dir := "/path/to/files"
	filesToZip, err := zipOps.ReadFiles(context.Background(), dir)
	if err != nil {
		fmt.Println("Failed to read files from directory:", err)
		return
	}
	zipBytes, err = zipOps.Zip(context.Background(), filesToZip)
	if err != nil {
		fmt.Println("Failed to zip files:", err)
		return
	}

	// Write data to a zip file
	err = zipOps.WriteZip(context.Background(), zipBytes, "archive.zip")
	if err != nil {
		fmt.Println("Failed to write zip file:", err)
		return
	}

	fmt.Println("Zip utility operations completed successfully!")
}

```
## Contributing

```
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

```

## License

```
This project is licensed under the MIT License.
Feel free to update the file with any additional information or formatting as needed.

```
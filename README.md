# Package Zip
The zip package provides a fast and simple zip operations.

- Zip files
- Unzip files
- Write files
- Read files

## Features
- One solution to all zip operations using golang

## Tech
```
- golang >= 1.17
```
## Installation
```
go get -u github.com/sdpsagarpawar/zip
```
## Uses

```sh
    zip := NewZipOperations()
	//from bytes to files
	files, err := zip.Unzip(context.Background(), []byte("your_zip_bytes"))
	//bytes to zipbytes
	zipBytes, err := zip.Zip(context.Background(), files)
	//from directory to file bytes
	fileBytes, err := zip.ReadFiles(context.Background(), "/your_directory")
	//from zipbytes to zip file
	err = zip.WriteZip(context.Background(), zipBytes, "expected_zip_name.zip")
```

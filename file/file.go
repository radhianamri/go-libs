package file

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

func IsExists(filepath string) bool {
	info, err := os.Stat(filepath)
	return err == nil && !info.IsDir()
}

func IsNotExists(filepath string) bool {
	return (!IsExists(filepath))
}

// List returns all files within a directory
func List(dirname string) ([]string, error) {
	var files []string
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

// Size returns the size of file without reading its content
func Size(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

// GetContent returns the whole content of a file
func GetContent(filename string) ([]byte, error) {
	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	contents, _ := io.ReadAll(bufio.NewReader(fp))
	return contents, nil
}

// PutContent writes the content into a file. It will attempt to create the file if it does not exist
func PutContent(filename string, content []byte) error {
	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(content)
	return err
}

// Copy copies the content of a source file to a destination file
func Copy(source, destination string) error {
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}

	return dest.Sync()
}

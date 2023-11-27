package file

import (
	"bufio"
	"io"
	"os"
)

func IsExists(filepath string) bool {
	info, err := os.Stat(filepath)
	return err == nil && !info.IsDir()
}

func IsNotExists(filepath string) bool {
	return (!IsExists(filepath))
}

func GetContents(filename string) ([]byte, error) {
	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)
	contents, _ := io.ReadAll(reader)
	return contents, nil
}

func PutContents(filename string, content []byte) error {
	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(content)
	return err
}

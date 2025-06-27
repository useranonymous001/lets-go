package storage

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// import (
// 	"fmt"
// 	"os"
// 	"bufio"
// 	"strings"
// )

type FileStorage struct {
	filename string
}

func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{filename: filename}
}

func (fs *FileStorage) Save(key, value string) error {

	// open/create a file
	file, err := os.OpenFile(fs.filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644) // file is type of io.Reader
	if err != nil {
		return err
	}
	// close the file at the end
	defer file.Close()

	// write to the file in key value pair format
	_, err = fmt.Fprintf(file, "%s=%s\n", key, value)

	// return err
	return err
}

func (fs *FileStorage) Load(key string) (string, error) {

	file, err := os.Open(fs.filename)
	if err != nil {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 && key == parts[0] {
			return parts[1], nil
		}
	}

	return "", fmt.Errorf("loading: key not found")

}

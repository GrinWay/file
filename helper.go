package file

import (
	"os"
	"fmt"
)

// read is an abstraction for real read work of certain readers
func read(filepath string, readable readerInterface) (any, error) {
	file, err := openFile(filepath)
	if nil != err {
		return nil, err
	}
	defer closeFile(file)
	return readable.Read(file)
}

func openFile(filepath string) (*os.File, error) {
	fmt.Printf("[INTERNAL] Opening file: \"%s\"\n", filepath)
	return os.Open(filepath)
}

func closeFile(file *os.File) error {
	fmt.Println("[INTERNAL] Closing file")
	return file.Close()
}

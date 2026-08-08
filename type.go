package file

import (
	"os"
)

// readerInterface interface is responsible how and which data to read from files
type readerInterface interface {
	// Supports answers able to support passed strategy or not
	Supports(strategy string) bool

	// Read is responsible to read all the necessary data from file according to the passed strategy
	Read(file *os.File) (any, error)
}

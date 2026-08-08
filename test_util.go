package file

import "os"

type customReader int

func (c customReader) Supports(strategy string) bool {
	return "custom" == strategy
}

func (c customReader) Read(file *os.File) (any, error) {
	return "test", nil
}

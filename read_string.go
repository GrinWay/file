package file

import (
	"os"
	"bufio"
)

type stringReader struct {}

func (r stringReader) Supports(strategy string) bool {
	return "strings" == strategy || "string" == strategy
}

func (r stringReader) Read(file *os.File) (any, error) {
	var textList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textList = append(textList, scanner.Text())
	}
	return textList, nil
}

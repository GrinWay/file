package file

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
)

type intReader struct {}

func (r intReader) Supports(strategy string) bool {
	return "ints" == strategy || "int" == strategy
}

func (r intReader) Read(file *os.File) (any, error) {
	var intList []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		intValue, err := strconv.ParseInt(text, 10, 64)
		if nil != err {
			fmt.Printf("[SKIP VALUE] \"%s\", not int\n", text)
			continue
		}
		intList = append(intList, int(intValue))
	}
	return intList, nil
}

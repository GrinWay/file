package file

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type numberReader struct{}

func (r numberReader) Supports(strategy string) bool {
	return "numbers" == strategy || "number" == strategy
}

func (r numberReader) Read(file *os.File) (any, error) {
	var numbers []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		number, err := strconv.ParseFloat(text, 64)
		if nil != err {
			fmt.Printf("[SKIP VALUE] \"%s\", not a number\n", text)
			continue
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

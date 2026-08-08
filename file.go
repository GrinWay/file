// Package file read segment data from file of a certain type
package file

import (
	"fmt"
)

var builtInReaders = []readerInterface{
	numberReader{},
	stringReader{},
	intReader{},
}

// Read accepts strategy, provided by existing structures
//  strategy := "my_own_reader_strategy_1"
//  filepath := os.Args[1]
//  values, err := file.Read(strategy, filepath, MyOwnReader1{}, MyOwnReader2{}, ...)
//  if nil != err {
//  	log.Fatal(err)
//  }
func Read(strategy string, filepath string, readers ...readerInterface) (any, error) {
	readers = append(readers, builtInReaders...)
	for _, reader := range readers {
		if !reader.Supports(strategy) {
			continue
		}
		return read(filepath, reader)
	}

	err := fmt.Errorf(
		"no supported reader for strategy: \"%s\", you can create your one",
		strategy,
	)
	return nil, err
}

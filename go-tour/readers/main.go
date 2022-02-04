package main

import (
	"errors"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	if len(b) == 0 && cap(b) == 0 {
		return 0, errors.New("err")
	}
	n := 0
	for i := range b {
		b[i] = 'A'
		n = i
	}
	return n + 1, nil
}
func main() {
	reader.Validate(MyReader{})
}

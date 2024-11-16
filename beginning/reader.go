package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (MyReader) Read(value []byte) (int, error) {
	for i := range value {
		value[i] = 'A'
	}
    return len(value), nil
}

func main() {
	reader.Validate(MyReader{})
}

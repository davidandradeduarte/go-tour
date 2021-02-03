package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func main() {
	reader.Validate(MyReader{})
}

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}

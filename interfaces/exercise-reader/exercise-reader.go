package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	n := copy(b, "A")
	return n, nil
}

func main() {
	reader.Validate(MyReader{})
	/*r := MyReader{}
	b := make([]byte, 8)
	n, err := r.Read(b)
	fmt.Printf("n = %v err = %v b = %v \n", n, err, b)*/
}

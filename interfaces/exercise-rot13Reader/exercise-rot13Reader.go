package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)

	for i, val := range b {
		if (val >= 65 && val <= 77) || (val >= 97 && val <= 109) {
			b[i] = val + 13
		} else if (val >= 78 && val <= 90) || (val >= 110 && val <= 122) {
			b[i] = val - 13
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

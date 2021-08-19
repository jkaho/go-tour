package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (int, error) {
	n, err := rr.r.Read(b)
	for i := 0; i < n; i++ {
		l := strings.ToLower(string(b[i]))
		if l >= "a" && l <= "m" {
			b[i] += 13
		} else if l > "m" && l <= "z" {
			b[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(r byte) byte {
	sb := rune(r)
	if sb >= 'a' && sb <= 'm' || sb >= 'A' && sb <= 'M' {
		sb = sb + 13
	} else if sb >= 'n' && sb <= 'z' || sb >= 'N' && sb <= 'Z' {
		sb = sb - 13
	}
	return byte(sb)

}

func (r13 *rot13Reader) Read(b []byte) (i int, e error) {
	n, err := r13.r.Read(b)
	for i := 0; i <= n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

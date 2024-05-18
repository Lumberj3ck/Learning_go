package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func rot13(char string) string {
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	uppwerCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	cases := [2]string{lowerCase, uppwerCase}

	for _, letterCase := range cases {
		if ind := strings.Index(letterCase, char); ind != -1 {
			new_ind := (ind + 13) % 26
			return string(letterCase[new_ind])
		}
	}
	return char
}

type rot13Reader struct {
	r io.Reader
}

type ReaderError string

func (e ReaderError) Error() string {
	return fmt.Sprintf("cannot read %e", string(e))
}

func (rot rot13Reader) Read(readed []byte) (int, error) {
	n, err := rot.r.Read(readed)
	if err != nil {
		return 0, ReaderError("Yee")
	}
	for i := 0; i < n; i++ {
		newLetter := rot13(string(readed[i]))
		readed[i] = newLetter[0]
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

package main

import (
	"fmt"
	"io"
	"math"
)

type Person struct {
	name struct {
		fname string
		lname string
	}
	age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v %v, aged %v", p.name.fname, p.name.lname, p.age)
}

// func main() {
// 	// p := Person{name: struct {
// 	// 	fname string
// 	// 	lname string
// 	// }{"Dennis", "Mithamo"}, age: 33}
// 	// fmt.Println(p)

// 	// _, err := sqrtWithErr(-10)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }

// 	// b := make([]byte, 8)
// 	// s := MyReader{}

// 	// for {
// 	// 	m, err := s.Read(b)

// 	// 	fmt.Printf("%s", b[:m])

// 	// 	if err == io.EOF {
// 	// 		break
// 	// 	}
// 	// }
// 	s := strings.NewReader("Lbh penpxrq gur pbqr!")
// 	r := rot13Reader{s}
// 	io.Copy(os.Stdout, &r)
// }

type ErrorUnsupportedComplex float64

func (e ErrorUnsupportedComplex) Error() string {
	return fmt.Sprintf("Cannot sqrt -ve number: %v", float64(e))
}

func sqrtWithErr(num float64) (float64, error) {
	if num < 0 {
		return 0, ErrorUnsupportedComplex(num)
	}

	guess := num / 2

	for ; math.Abs(guess*guess-num) > 0.00001; guess -= (guess*guess - num) / (2 * guess) {
	}

	return guess, nil
}

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	lenStream, err := rot.r.Read(b)
	if err != nil {
		return lenStream, err
	}

	for i := 0; i < lenStream; i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = 'A' + (b[i]-'A'+13)%26
		} else if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = 'a' + (b[i]-'a'+13)%26
		}
	}

	return lenStream, nil
}

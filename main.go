package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"unicode"
	"log"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(ScanFrame)

	for r.Scan() {
	}
}

func ScanFrame(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	buf := bytes.NewBuffer(data)
outer:
	for {
		start++
		rune, _, err := buf.ReadRune()
		if err == io.EOF {
			break
		}
		if rune != '\033' {
			continue
		}

		start++
		rune, _, err = buf.ReadRune()
		if err == io.EOF {
			break
		}
		if rune != '[' {
			continue
		}

		x := bytes.NewBufferString("")
		for {
			start++
			rune, _, err = buf.ReadRune()
			if err == io.EOF {
				break outer
			}
			if !unicode.IsDigit(rune) {
				buf.UnreadRune()
				break
			}
			x.WriteRune(rune)
		}
		if x.Len() == 0 {
			continue
		}
		xInt, err := strconv.Atoi(x.String())
		if err != nil {
			continue
		}

		start++
		rune, _, err = buf.ReadRune()
		if err == io.EOF {
			break
		}
		if rune != ';' {
			continue
		}

		y := bytes.NewBufferString("")
		for {
			start++
			rune, _, err = buf.ReadRune()
			if err == io.EOF {
				break outer
			}
			if !unicode.IsDigit(rune) {
				buf.UnreadRune()
				break
			}
			y.WriteRune(rune)
		}
		if y.Len() == 0 {
			continue
		}
		yInt, err := strconv.Atoi(y.String())
		if err != nil {
			continue
		}

		start++
		rune, _, err = buf.ReadRune()
		if err == io.EOF {
			break
		}
		if rune != 'f' {
			continue
		}
		log.Printf("found %d %d", xInt, yInt)
	}
	return start, nil, nil
}

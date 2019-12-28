package main

import (
	"bufio"
	"os"
	"time"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(ScanFrame)

	for r.Scan() {
		os.Stdout.Write(r.Bytes())
		//fmt.Printf("%x", r.Text())
		time.After(time.Second)

	}
}

func ScanFrame(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	for ; start < len(data)-4; start++ {
		if string(data[start:start+4]) == "\033[0m" {
			return start+5, data[:start+4], nil
		}
	}
	return start, nil, nil
}

package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	c := mqtt.NewClient(mqtt.NewClientOptions().AddBroker("tcp://mqtt.iot.exposed:1883"))
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanRunes)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	frame := bytes.NewBufferString("")
	for r.Scan() {
		_, _ = frame.Write(r.Bytes())
		if strings.HasSuffix(frame.String(), "\033[0") {
			os.Stdout.Write(frame.Bytes())
			os.Stdout.Sync()
			token := c.Publish("text", 0, false, "\r" + frame.String())
			token.Wait()
			frame = bytes.NewBufferString("")
		}
	}
}

//func ScanFrame(data []byte, atEOF bool) (advance int, token []byte, err error) {
//	start := 0
//	for ; start < len(data)-4; start++ {
//		if string(data[start:start+4]) == "\033[0m" {
//			return start + 4, data[:start+4], nil
//		}
//	}
//	return start, nil, nil
//}

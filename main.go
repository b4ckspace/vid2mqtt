package main

import (
	"bufio"
	"bytes"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	c := mqtt.NewClient(mqtt.NewClientOptions().AddBroker("tcp://mqtt.iot.exposed:1883"))
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanRunes)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	token := make([]byte, 3)
	frame := bytes.NewBufferString("")
	for r.Scan() {
		b := r.Bytes()
		_, _ = frame.Write(b)
		token[0] = token[1]
		token[1] = token[2]
		token[2] = b[0]
		if string(token) == "\033[0" {
			os.Stdout.WriteString("\r" + frame.String())
			os.Stdout.Sync()
			token := c.Publish("text", 0, false, "\r"+frame.String())
			token.Wait()
			frame.Reset()
		}
	}
}

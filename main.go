package main

import (
	"bufio"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	c := mqtt.NewClient(mqtt.NewClientOptions().AddBroker("tcp://mqtt.iot.exposed:1883"))
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Mqtt error: %s", token.Error())
	}

	s := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	s.Buffer(buf, 1024*1024)

	for s.Scan() {
		text := s.Text()
		os.Stdout.WriteString(text)
		token := c.Publish("text", 0, false, "\r"+text)
		token.Wait()
	}

	log.Fatalf("Scan aborted: %s", s.Err())
}

func ScanFrame(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	for ; start < len(data)-3; start++ {
		if string(data[start:start+3]) == "\033[0" {
			return start + 3, data[:start+3], nil
		}
	}
	return 0, nil, nil
}

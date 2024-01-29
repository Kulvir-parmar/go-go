package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type server struct {
	msgChannel  chan Message
	quitChannel chan struct{}
}

func (s *server) init() {
	// this is a name for the **for** loop
free:
	for {
		select {
		case msg := <-s.msgChannel:
			fmt.Printf("message from %s: and %s \n", msg.From, msg.Payload)
		case <-s.quitChannel:
			fmt.Println("quitting from the channel")
			// logic to close the channel
			break free
		default:
			continue
		}
	}
	fmt.Println("the server is down")
}

func quitChannel(channel chan struct{}) {
	close(channel)
}

func sendMessage(msgChannel chan Message, payload string) {
	msg := Message{
		From:    "JJ",
		Payload: payload,
	}

	msgChannel <- msg
}

func main() {
	s := server{
		msgChannel:  make(chan Message),
		quitChannel: make(chan struct{}),
	}

	go s.init()

	go func() {
		time.Sleep(2 * time.Second)
		sendMessage(s.msgChannel, "Hello World")
	}()

	go func() {
		time.Sleep(4 * time.Second)
		quitChannel(s.quitChannel)
	}()

	select {}
}

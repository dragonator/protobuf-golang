package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"google.golang.org/protobuf/proto"

	"github.com/dragonator/protobuf-golang/internal/pb/messages"
)

func main() {
	fmt.Println("Starting client..")
	m := &messages.Message{
		Sender: "Bay Ivan",
		Text:   "Hello!",
	}
	message, err := proto.Marshal(m)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while marshaling message: %s", err.Error())
		os.Exit(1)
	}

	_, err = http.Post("http://localhost:8080/protobuf", "", bytes.NewBuffer(message))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while post request to server: %s", err.Error())
		os.Exit(1)
	}
	_, err = http.Post("http://localhost:8080/message", "", bytes.NewBuffer(message))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while post request to server: %s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Sent %d bytes to server\n", len(message))
}

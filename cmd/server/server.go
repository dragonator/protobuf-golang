package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/protobuf/proto"

	"github.com/dragonator/protobuf-golang/internal/pb/messages"
)

func main() {
	http.HandleFunc("/protobuf", ProtoBufHandler)
	http.HandleFunc("/message", MessageHandler)

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ProtoBufHandler -
func ProtoBufHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Bytes: %v\n", data)
	fmt.Printf("Base64: %s\n", base64.StdEncoding.EncodeToString(data))
	fmt.Printf("HEX: %s\n", hex.EncodeToString(data))
}

// MessageHandler -
func MessageHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	m := messages.Message{}
	if err := proto.Unmarshal(data, &m); err != nil {
		fmt.Println(err)
	}

	fmt.Println(m.String())
}

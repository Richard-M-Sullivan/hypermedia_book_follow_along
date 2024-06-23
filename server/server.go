package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func PrintServer() {
	fmt.Println("Server started!")
}

func StartServer() {
	// create handler for the server
	http.HandleFunc("/", helloHandler)

	// log that the server stopped when this exits
	PrintServer()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World\n")
}

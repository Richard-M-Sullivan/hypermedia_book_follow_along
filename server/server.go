package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func PrintServer() {
	fmt.Println("server package imported")
}

func StartServer() {
	// create handler for the server
	http.HandleFunc("/", helloHandler)

	// log that the server stopped when this exits
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World\n")
}

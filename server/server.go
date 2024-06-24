package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func printServer() {
	fmt.Println("Server started!")
}

func StartServer() {
	// create handler for the server
	mux := http.NewServeMux()

	// adding handlers here
	mux.Handle("/", http.HandlerFunc(rootHandler))
	mux.Handle("/contacts", http.HandlerFunc(contactsHandler))

	// log that the server stopped when this exits
	printServer()
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// handle the root request by redirecting clients to the contacts page
func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("rootHandler:", req.URL)
	http.Redirect(w, req, "contacts", http.StatusMovedPermanently)
}

func contactsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("contactsHandler", req.URL)
	io.WriteString(w, "Contacts Page\n")
}

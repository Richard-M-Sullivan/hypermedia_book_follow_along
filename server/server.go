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

// launch contacts website server
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
	http.Redirect(w, req, "/contacts", http.StatusMovedPermanently)
}

// return contacts page
func contactsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("contactsHandler", req.URL)

	io.WriteString(w, "Contacts Page\n")

	// if q not found show all contacts, otherwise filter the search with the query
	req.ParseForm()

	if req.Form.Has("q") {
		io.WriteString(w, fmt.Sprintf("Searching for %v in contacts\n", req.Form["q"]))
	} else {
		io.WriteString(w, "Showing all contacts\n")
	}
}

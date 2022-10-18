package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	_myhandler := func(_writer http.ResponseWriter, _request *http.Request) {
		io.WriteString(_writer, "Check state!\n")
	}

	http.HandleFunc("/hello", _myhandler)
	log.Println("Listen for request!")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

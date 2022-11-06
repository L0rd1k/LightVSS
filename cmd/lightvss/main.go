package main

import (
	"fmt"
	"log"
	"net/http"
)

func test_handler(_writer http.ResponseWriter, _request *http.Request) {
	fmt.Fprintf(_writer, "Jopa %s", _request.URL.Path)
}

func test_handler_2(_writer http.ResponseWriter, _request *http.Request) {
	for k, v := range _request.Header {
		fmt.Fprintf(_writer, "Header[%q] = %q\n", k, v)
	}
}

func main() {

	// _myhandler := func(_writer http.ResponseWriter, _request *http.Request) {
	// 	io.WriteString(_writer, "Check state!\n")
	// }
	// http.HandleFunc("/hello", _myhandler)
	// log.Println("Listen for request!")

	http.HandleFunc("/", test_handler)
	http.HandleFunc("/hello", test_handler_2)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

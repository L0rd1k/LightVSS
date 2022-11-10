package main

import (
	"fmt"
	"net/http"

	"github.com/L0rd1k/LightVSS/test/containers/net"
)

func test_handler(_writer http.ResponseWriter, _request *http.Request) {
	fmt.Fprintf(_writer, "Jopa %s", _request.URL.Path)
}

func test_handler_2(_writer http.ResponseWriter, _request *http.Request) {
	for k, v := range _request.Header {
		fmt.Fprintf(_writer, "Header[%q] = %q\n", k, v)
	}
}

type Handler struct{}

func (handler *Handler) ServeHTTP(_writer http.ResponseWriter, _request *http.Request) {
	switch _request.URL.Path {
	case "/":
		fmt.Fprintf(_writer, "URL.Path = %q\n", _request.URL.Path)
	case "/hello":
		for k, v := range _request.Header {
			fmt.Fprintf(_writer, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(_writer, "404 NOT FOUND: %s\n", _request.URL)
	}
}

func main() {

	reader := net.New()

	reader.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	reader.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	reader.Run(":9999")
}

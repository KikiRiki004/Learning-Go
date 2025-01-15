package main

import (
	"fmt"
	"net/http" // package for basic server
)

func hello(w http.ResponseWriter, req *http.Request) { // A handler is an object implementing the http.Handler interface.

	// Functions serving as handlers take a http.ResponseWriter and a http.Request as arguments.
	// The response writer is used to fill in the HTTP response.

	fmt.Fprintf(w, "hello\n") // Simple response
}

func headers(w http.ResponseWriter, req *http.Request) {
	// Reading all the HTTP request headers and echoing them into the response body.

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// We register our handlers on server routes using the http.HandleFunc convenience function.
	// It sets up the default router in the net/http package and takes a function as an argument.

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil) // nil tells it to use the default router set up above
}

//Run the server in the background:
//$ go run http-servers.go &
//Access the /hello route:
//$ curl localhost:8090/hello
//output: hello

package main

import (
	"fmt"
	"net/http"
	"time"
)

// A Context carries deadlines, cancellation signals, and other request-scoped values across API boundaries and goroutines.

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

// $ go run context.go

// $ curl localhost:8090/hello
// server: hello handler started
// ^C
// server: context canceled
// server: hello handler ended

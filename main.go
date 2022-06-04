package main

import (
	"fmt"
	"http-hello-world/utils"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, utils.GetPrefix()+"Hello, you've requested: %s\n", request.URL.Path)
	})

	http.ListenAndServe(":3000", nil)
}

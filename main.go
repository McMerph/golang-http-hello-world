package main

import (
	"fmt"
	"http-hello-world/utils"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, utils.GetPrefix()+"You've requested: %s\n", request.URL.Path)
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

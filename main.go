package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("start server ...")
	startHttpServer()
}

func startHttpServer() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "hail hydra")
	})
	http.ListenAndServe(":9090", nil)
}

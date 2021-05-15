package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start server ...")
	startHttpServer()
}

func startHttpServer() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "hail hydra")
		log.Println("recv: "+ r.RequestURI)
	})
	http.ListenAndServe(":9090", nil)
}

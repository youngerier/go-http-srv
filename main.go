package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	_ = os.Mkdir("./log",os.ModePerm)
	file, err := os.OpenFile("./log/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	
	log.Println("Hello world!")
}

func main() {
	fmt.Println("start server ...")
	startHttpServer()
}

func startHttpServer() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "hail hydra")
		log.Println("recv: " + r.RequestURI)
	})
	http.ListenAndServe(":9090", nil)
}

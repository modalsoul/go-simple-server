package main

import (
	"net/http"
	. "modalSoul"
	"flag"
	"fmt"
)

func main() {
	f := flag.Int("port", 3000, "port 3000")
	flag.Parse()
	serve := fmt.Sprintf(":%d", *f)
	http.HandleFunc("/", Handler)
	http.ListenAndServe(serve, nil)
}
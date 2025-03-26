package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

func main() {
	var dir, addr string
	flag.StringVar(&dir, "dir", "../dist", "dir to serve")
	flag.StringVar(&addr, "addr", "0.0.0.0:8080", "addr to bind")
	flag.Parse()

	s := &http.Server{
		Addr:           addr,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.Handle("/", http.FileServer(http.Dir(dir)))

	log.Printf("Serving files from %s on %s\n", dir, addr)
	log.Fatal(s.ListenAndServe())
}

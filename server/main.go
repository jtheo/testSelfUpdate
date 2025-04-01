package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

type dirpath string

func (p dirpath) serveFile(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fullpath := path.Join(string(p), r.RequestURI)
	status := http.StatusOK
	_, err := os.Stat(fullpath)
	if err != nil {
		if os.IsNotExist(err) {
			status = http.StatusNotFound
			log.Printf("%3d % 10s %s \n", status, time.Since(start).Round(time.Nanosecond), fullpath)
			http.NotFound(w, r)
			return
		}
		status = http.StatusInternalServerError
		fmt.Fprint(w, "Internal Server Error")
		log.Printf("%3d % 10s %s \n", status, time.Since(start).Round(time.Nanosecond), fullpath)
		return
	}
	http.ServeFile(w, r, fullpath)
	log.Printf("%3d % 10s %s \n", status, time.Since(start).Round(time.Nanosecond), fullpath)
}

func main() {
	var dir, addr string
	flag.StringVar(&dir, "dir", "../dist", "dir to serve")
	flag.StringVar(&addr, "addr", "0.0.0.0:8080", "addr to bind")
	flag.Parse()
	p := dirpath(dir)

	s := &http.Server{
		Addr:           addr,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/", p.serveFile)

	log.Printf("Serving files from %s on %s\n", dir, addr)
	log.Fatal(s.ListenAndServe())
}

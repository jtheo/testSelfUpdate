package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/minio/selfupdate"
)

var Version string = "0"

func main() {
	var showVer bool
	var baseUrl string
	flag.BoolVar(&showVer, "V", false, "show version and exit")
	flag.StringVar(&baseUrl, "b", "https://hako.us/self", "base url")
	flag.Parse()

	if showVer {
		fmt.Println("Version:", Version)
		return
	}

	if err := doUpdate(baseUrl); err != nil {
		log.Fatalln(err)
	}
}

func doUpdate(url string) error {
	url = url + "/" + runtime.GOOS
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = selfupdate.Apply(resp.Body, selfupdate.Options{})
	if err != nil {
		if rerr := selfupdate.RollbackError(err); rerr != nil {
			fmt.Printf("Failed to rollback from bad update: %v", rerr)
		}
	}
	return err
}

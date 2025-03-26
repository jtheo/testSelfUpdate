package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/minio/selfupdate"
)

var Version string = "0"

func main() {
	start()
	fmt.Println("Hello, World!")
}

func start() {
	var showVer, update bool
	var baseURL string
	flag.StringVar(&baseURL, "b", "http://localhost:8080", "base url")
	flag.BoolVar(&update, "u", false, "selftupdate")
	flag.BoolVar(&showVer, "v", false, "Show version and exit")
	flag.Parse()

	if update {
		if checkIsUpdated(baseURL) {
			fmt.Println("You are already at the latest version")
			os.Exit(0)
		}
		if err := doUpdate(baseURL); err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}
	checkIsUpdated(baseURL)

	if showVer {
		fmt.Printf("%s, version: %s\n", path.Base(os.Args[0]), Version)
		os.Exit(0)
	}
}

func doUpdate(url string) error {
	fmt.Println("Running the update")
	myself := path.Base(os.Args[0])
	version := getVersion(url)
	url = url + "/" + version + "/" + myself + "-" + runtime.GOOS + "-" + runtime.GOARCH
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("the HTTP requests for %s returned %d\n", url, resp.StatusCode)
	}
	err = selfupdate.Apply(resp.Body, selfupdate.Options{})
	if err != nil {
		log.Printf("Something is wrong: \n%v\n", err)
		if rerr := selfupdate.RollbackError(err); rerr != nil {
			fmt.Printf("Failed to rollback from bad update: %v", rerr)
		}
	}
	fmt.Println("Update successful, bye!")
	return nil
}

func getVersion(url string) string {
	url = url + "/version"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	version := strings.Trim(string(body), "\n")
	return version
}

func checkIsUpdated(url string) bool {
	newVersion := getVersion(url)

	old, err := strconv.Atoi(Version)
	if err != nil {
		panic(err)
	}
	new, err := strconv.Atoi(newVersion)
	if err != nil {
		panic(err)
	}

	if old < new {
		fmt.Printf("Your version of %s is out of date\nYou are on version %s, but there's a new version %s\n\n", path.Base(os.Args[0]), Version, newVersion)
		return false
	}
	return true
}

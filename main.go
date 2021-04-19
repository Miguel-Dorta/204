package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var version = "Unknown"

func init() {
	showVersion := false
	flag.BoolVar(&showVersion, "v", false, "Show version and exit")
	flag.BoolVar(&showVersion, "version", false, "Show version and exit")
	flag.Parse()
	if showVersion {
		fmt.Printf("Project: 204\nVersion: %s\nAuthor:  Miguel Dorta <contact@migueldorta.com>\nSource:  github.com/Miguel-Dorta/204\n", version)
		os.Exit(0)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		if _, err := w.Write([]byte{}); err != nil {
			log.Printf("Error writing response: %s\n", err.Error())
			return
		}
	})

	if err := new(http.Server).ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Unexpected error closed the server: %s", err.Error())
		return
	}
}

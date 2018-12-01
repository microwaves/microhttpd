package main

import (
	"flag"
	"log"
	"net/http"
)

var port string
var staticPath string

func init() {
	flag.StringVar(&port, "port", "80", "port to run the httpd")
	flag.StringVar(&staticPath, "serve", "www/", "path to the files to be served")
	flag.Parse()
}

func main() {
	fs := http.FileServer(http.Dir(staticPath))
	http.Handle("/", http.StripPrefix("/", fs))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

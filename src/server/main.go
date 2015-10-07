package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := fmt.Sprintf("%s:%s",
		os.Getenv("WEBSERVERHOST"), os.Getenv("WEBSERVERPORT"))

	log.Fatal(http.ListenAndServe(addr), NewRouter())
}

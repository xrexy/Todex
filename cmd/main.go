package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/xrexy/Todex/src/pkg/http/rest"
)

func main() {
	port := flag.Int("port", 8080, "port to run the server on")
	flag.Parse()

	fmt.Println("Starting server on port:", *port)

	router := rest.InitHandlers()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}

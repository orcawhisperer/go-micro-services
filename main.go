package main

import (
	"log"
	"net/http"
	"os"

	"github.com/iamvasanth07/go-micro-services/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":8080", sm)
}

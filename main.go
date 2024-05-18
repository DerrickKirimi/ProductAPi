package main

import (
	"log"
	"net/http"
	"os"

	"github.com/DerrickKirimi/ProductAPi/handlers"
)

func main() {
	//Handlefunc adds func to defaultServeMux
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	h := handlers.NewHello(l)
	g := handlers.NewGoodbye(l)
	sm :=  http.NewServeMux()
	sm.Handle("/", h)
	sm.Handle("/goodbye", g)
	http.ListenAndServe(":9090", sm)
}
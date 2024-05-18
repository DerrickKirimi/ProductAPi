package main

import (
	"log"
	"net/http"
	"os"
	"time"

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

	s := &http.Server{
		Addr:			":9090",
		Handler:		sm,
		IdleTimeout:	120 * time.Second,
		ReadTimeout:	1 * time.Second,
		WriteTimeout:	1 * time.Second,
	}

	s.ListenAndServe()
}
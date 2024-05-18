package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	go func () {
	err := s.ListenAndServe()
	if err != nil {
		l.Fatal(err)
	}
	} ()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <- sigChan
	l.Println("Received terminate signal, graceful shutdown", sig)

	ct,_ := context.WithTimeout(context.Background(), 30 * time.Second)
	s.Shutdown(ct)


}
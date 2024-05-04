package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//Handlefunc adds func to defaultServeMux
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request) {
		log.Println("Ssup G")
		d, _ := io.ReadAll(r.Body)
		fmt.Fprintf(rw, "Hello %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request){
		log.Println("Aloha")
	})
	http.ListenAndServe(":9090", nil)
}

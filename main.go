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
		d, err := io.ReadAll(r.Body)
		if err != nil{
			http.Error(rw, "Oops", http.StatusBadRequest)
			return //Needed as Error does not end the request
		}
		fmt.Fprintf(rw, "Hello %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request){
		log.Println("Aloha")
	})
	http.ListenAndServe(":9090", nil)
}

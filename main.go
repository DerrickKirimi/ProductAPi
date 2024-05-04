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
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Ooops"))
			return
		}
		fmt.Fprintf(rw, "Hello %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request){
		log.Println("Aloha")
	})
	http.ListenAndServe(":9090", nil)
}

package main

import (
	"log"
	"net/http"
)

func main() {
	//Handlefunc adds func to defaultServeMux
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Ssup G")
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request){
		log.Println("Aloha")
	})
	http.ListenAndServe(":9090", nil)
}

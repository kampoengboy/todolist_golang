package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

	// router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	log.Fatal(http.ListenAndServe(":8080", router))
}

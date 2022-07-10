package main

import (
	endpoint "example/keystore/endpoints"
	"log"
	"net/http"
)

func main() {

	keystore := endpoint.NewstoreHandlers()

	http.HandleFunc("/get/", keystore.Getkeys)
	http.HandleFunc("/search", keystore.Searchkeys)
	http.HandleFunc("/set", keystore.Setkeys)
	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
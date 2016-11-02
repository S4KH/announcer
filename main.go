/**
* Entry point of the system
*
* @author skh
**/

package main

import (
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(port, router))
}

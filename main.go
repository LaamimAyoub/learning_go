package main

import (
	"log"
	"mymodule/dep"
	"net/http"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	dep.Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}

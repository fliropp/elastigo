package main

import (
	"fmt"
	"net/http"
)

func main() {

	/* init controllers */
	fmt.Printf("init controllers...")
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/ping", pingHandler)

	http.ListenAndServe(":8080", nil)

}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}

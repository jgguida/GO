package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/saudacao/", func(w http.ResponseWriter, r *http.Request) {
		nome := r.URL.Path[len("/saudacao/"):]
		fmt.Fprintf(w, "Olá, %s!", nome)
	})

	http.ListenAndServe(":8080", nil)
}

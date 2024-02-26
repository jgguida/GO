package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/texto-mockado", textoMockadoHandler)
	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}

func textoMockadoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "Este é um texto mockado retornado pela API!")
}

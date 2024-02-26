package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/saudacao", saudacaoHandler)

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("form").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Formulário de Saudação</title>
		</head>
		<body>
			<form action="/saudacao" method="post">
				<label for="nome">Digite seu nome:</label>
				<input type="text" id="nome" name="nome">
				<input type="submit" value="Enviar">
			</form>
		</body>
		</html>
	`))
	tmpl.Execute(w, nil)
}

func saudacaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	nome := r.FormValue("nome")
	fmt.Fprintf(w, "Olá, %s!", nome)                // Escreve na tela
	w.Header().Set("X-Custom-Header", "Olá, "+nome) // Define cabeçalho personalizado
}

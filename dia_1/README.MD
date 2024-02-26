# Servidor HTTP Básico
Este guia explicará cada linha do código fornecido, o qual implementa um servidor HTTP simples em Go que responde a uma rota específica com um texto mockado. Este exemplo é uma ótima maneira de começar a entender como criar aplicações web básicas com Go.

Estrutura do Código
go
Copy code
package main
package main: Define o pacote ao qual este arquivo pertence. Em Go, o pacote main é especial: ele define um executável, não uma biblioteca.
go
Copy code
import (
	"fmt"
	"net/http"
)
import: Esta seção importa os pacotes necessários para o nosso código.
fmt: Utilizado para operações de entrada/saída formatadas, como imprimir mensagens no console.
net/http: Um pacote poderoso que permite a implementação de clientes e servidores HTTP em Go.
go
Copy code
func main() {
	http.HandleFunc("/texto-mockado", textoMockadoHandler)
	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
func main(): A função principal do nosso programa. É o ponto de entrada do executável.
http.HandleFunc("/texto-mockado", textoMockadoHandler): Registra a função textoMockadoHandler como o manipulador para a rota /texto-mockado. Quando essa rota for acessada, a função textoMockadoHandler será chamada.
fmt.Println("Servidor rodando na porta 8080..."): Imprime uma mensagem no console indicando que o servidor está rodando e em qual porta.
http.ListenAndServe(":8080", nil): Inicia um servidor HTTP na porta 8080. O segundo argumento (nil) indica que estamos usando o multiplexador de rotas padrão.
go
Copy code
func textoMockadoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "Este é um texto mockado retornado pela API!")
}
func textoMockadoHandler(w http.ResponseWriter, r *http.Request): Define o manipulador para a rota /texto-mockado.
w http.ResponseWriter: Um objeto que representa a resposta que será enviada ao cliente.
r *http.Request: Um ponteiro para o objeto que representa a requisição recebida do cliente.
if r.Method != http.MethodGet: Verifica se o método da requisição é GET. Se não for, retorna um erro 405 Method Not Allowed.
http.Error(w, "Método não suportado", http.StatusMethodNotAllowed): Envia uma resposta de erro ao cliente com a mensagem "Método não suportado".
fmt.Fprint(w, "Este é um texto mockado retornado pela API!"): Se a requisição for um GET, escreve a mensagem "Este é um texto mockado retornado pela API!" na resposta.
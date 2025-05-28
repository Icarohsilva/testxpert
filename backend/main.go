package main

import (
	"encoding/json" // pra transformar dados em JSON
	"fmt"           // pra mostrar mensagens no terminal
	"net/http"      // pra criar um servidor web
)

type TestCase struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var testCases = []TestCase{
	{ID: 1, Title: "Login com sucesso", Description: "Validar login com credenciais corretas"},
	{ID: 2, Title: "Login com falha", Description: "Validar login com credenciais incorretas"},
	{ID: 3, Title: "Cadastro de usuário", Description: "Validar cadastro de um novo usuário"},
}

func main() {
	http.HandleFunc("/testcases", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(testCases)
	})

	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}

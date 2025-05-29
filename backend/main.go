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

type BugReports struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

var testCases = []TestCase{
	{ID: 1, Title: "Login com sucesso", Description: "Validar login com credenciais corretas"},
	{ID: 2, Title: "Login com falha", Description: "Validar login com credenciais incorretas"},
	{ID: 3, Title: "Cadastro de usuário", Description: "Validar cadastro de um novo usuário"},
	{ID: 4, Title: "Edição de perfil", Description: "Validar edição de dados do perfil do usuário"},
	{ID: 5, Title: "Redefinição de senha", Description: "Validar processo de redefinição de senha"},
}

var bugReports = []BugReports{}

var bugId = 1

// Função principal que inicia o servidor HTTP

func main() {
	http.HandleFunc("/testcases", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") // Define o tipo de resposta como JSON
		json.NewEncoder(w).Encode(testCases)               // Converte a lista em JSON e envia
	})

	// Rota para criar um novo bug report
	http.HandleFunc("/bugs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodGet {
			var newBug BugReports
			json.NewDecoder(r.Body).Decode(&newBug)
			newBug.ID = bugId
			bugId++
			bugReports = append(bugReports, newBug)

			// Responde com o novo bug criado
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newBug)
		} else {
			json.NewEncoder(w).Encode(bugReports) // Retorna todos os bugs existentes
		}
	})

	fmt.Println("TestXpert rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

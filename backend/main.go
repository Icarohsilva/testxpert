package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Estrutura de um caso de teste
type TestCase struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Estrutura de um bug reportado
type BugReport struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

// Lista de casos de teste iniciais
var testCases = []TestCase{
	{ID: 1, Title: "Login com sucesso", Description: "Validar login com credenciais válidas"},
	{ID: 2, Title: "Senha inválida", Description: "Exibir erro ao inserir senha incorreta"},
	{ID: 3, Title: "Campo obrigatório", Description: "Validar mensagem de erro ao deixar campo obrigatório vazio"},
	{ID: 4, Title: "Redirecionamento após login", Description: "Verificar redirecionamento para a página inicial após login bem-sucedido"},
	{ID: 5, Title: "Logout", Description: "Validar que o usuário é desconectado corretamente"},
	{ID: 6, Title: "Cadastro de usuário", Description: "Validar cadastro com dados válidos"},
}

// Lista de bugs (inicialmente vazia)
var bugReports = []BugReport{}
var bugID = 1

func main() {
	// Endpoint para retornar casos de teste
	http.HandleFunc("/testcases", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(testCases)
	})

	// Endpoint para manipular bugs
	http.HandleFunc("/bugs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodPost {
			var newBug BugReport
			json.NewDecoder(r.Body).Decode(&newBug)

			// Validação simples para garantir que o título e os detalhes não estejam vazios
			if newBug.Title == "" || newBug.Details == "" {
				http.Error(w, "Título e detalhes são obrigatórios", http.StatusBadRequest)
				return
			}
			newBug.ID = bugID
			bugID++
			bugReports = append(bugReports, newBug)

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newBug)

		} else {
			json.NewEncoder(w).Encode(bugReports)
		}
	})

	fmt.Println("TestXpert rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

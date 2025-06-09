package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Structs já existentes

type TestCase struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BugReport struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

var testCases = []TestCase{
	{ID: 1, Title: "Login com sucesso", Description: "Validar login com credenciais válidas"},
	{ID: 2, Title: "Senha inválida", Description: "Exibir erro ao inserir senha incorreta"},
	{ID: 3, Title: "Campo obrigatório", Description: "Validar que o campo 'nome' é obrigatório"},
	{ID: 4, Title: "Redirecionamento após login", Description: "Verificar redirecionamento para a página inicial após login bem-sucedido"},
	{ID: 5, Title: "Logout", Description: "Validar que o usuário pode fazer logout com sucesso"},
	{ID: 6, Title: "Validação de email", Description: "Verificar se o email inserido é válido"},
}

var bugReports = []BugReport{}
var bugID = 1

// INÍCIO DO TRECHO NOVO DO EPISÓDIO 5
type TestPlan struct {
	Features string `json:"features"`
	TestCases []TestCase `json:"test_cases"`
	Equipe string `json:"testers"`
	Responsavel string `json:"responsible"`
	Artefatos []string `json:"artifacts"`
}

var testPlan = TestPlan{
	Features: "Funcionalidade de login no sistema",
	TestCases: testCases,
	Equipe: "Equipe de QA",
	Responsavel: "João da Silva",
	Artefatos: []string{"Documentação de requisitos", "Casos de teste", "Relatórios de bugs"},
}	

// FIM DO TRECHO NOVO DO EPISÓDIO 5

func main() {
	http.HandleFunc("/testcases", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(testCases)
	})

	// INÍCIO DO TRECHO NOVO DO EPISÓDIO 5
	// Endpoint para retornar o plano de testes
	http.HandleFunc("/testplan", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodGet {
			json.NewEncoder(w).Encode(testPlan)
		}
	})
			
	// FIM DO TRECHO NOVO DO EPISÓDIO 5

	http.HandleFunc("/bugs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodPost {
			var newBug BugReport
			json.NewDecoder(r.Body).Decode(&newBug)

			if newBug.Title == "" {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "O campo 'title' é obrigatório."})
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

	// Inicia o servidor
	fmt.Println("TestXpert rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
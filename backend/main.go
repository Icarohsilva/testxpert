package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Structs já existentes

type TestCase struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority   	string `json:"priority,omitempty"` // Campo opcional
	Risk		string `json:"risk,omitempty"`     // Campo opcional
}

type BugReport struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

type TestPlan struct {
	Features string `json:"features"`
	TestCases []TestCase `json:"test_cases"`
	Equipe string `json:"testers"`
	Responsavel string `json:"responsible"`
	Artefatos []string `json:"artifacts"`
	Priority    string  `json:"priority"`  // Novo campo
}


var testCases = []TestCase{
	{ID: 1, Title: "Login com sucesso - obrigatório", Description: "Validar login com credenciais válidas"},
	{ID: 2, Title: "Login com senha inválida", Description: "Validar mensagem de erro ao tentar logar com senha inválida", Risk: "Alto"},
	{ID: 3, Title: "Login com campo obrigatório vazio", Description: "Validar mensagem de erro ao deixar campo obrigatório vazio", Risk: "Médio"},
	{ID: 4, Title: "Login com usuário inexistente", Description: "Validar mensagem de erro ao tentar logar com usuário inexistente"},
	{ID: 5, Title: "Login com sessão expirada", Description: "Validar redirecionamento para a página de login ao tentar acessar uma página com sessão expirada"},
	{ID: 6, Title: "Login com múltiplas tentativas", Description: "Validar bloqueio de conta após múltiplas tentativas de login com senha inválida - crítico",Risk: "Médio"},
	{ID: 7, Title: "Login com autenticação de dois fatores", Description: "Validar fluxo de autenticação de dois fatores"},

}

var bugReports = []BugReport{}
var bugID = 1

var testPlan = TestPlan{
	Features: "Funcionalidade de login no sistema",
	TestCases: testCases,
	Equipe: "Equipe de QA",
	Responsavel: "João da Silva",
	Artefatos: []string{"Documentação de requisitos", "Casos de teste", "Relatórios de bugs"},
	Priority: "Alta", // Definindo a prioridade do plano de testes
	
}	

func (tc *TestCase) DeterminePriority() string {
	if tc.Risk == "Alto" {
		return "Alto"
	}
	if strings.Contains(tc.Title, "obrigatório") || strings.Contains(tc.Title, "inválida") {
		return "Alta"
	}
	if strings.Contains(tc.Description, "crítico") || strings.Contains(tc.Description, "segurança") {
		return "Alta"
	}
	return "Média"
	
}

func main() {
	http.HandleFunc("/testcases", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(testCases)
	})

	// Endpoint para retornar o plano de testes
	http.HandleFunc("/testplan", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodGet {
			json.NewEncoder(w).Encode(testPlan)
		}
	})

	// INÍCIO DO TRECHO NOVO DO EPISÓDIO 6
	// Endpoint para retornar casos de teste de alta prioridade
	http.HandleFunc("/priority-tests", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		for i := range testCases {
			testCases[i].Priority = testCases[i].DeterminePriority()
		}
		higtPriorityTests := []TestCase{}
		for _, tc := range testCases {
			if tc.Priority == "Alta" || tc.Priority == "Alto" {
				higtPriorityTests = append(higtPriorityTests, tc)
			}
		}
		if len(higtPriorityTests) == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Nenhum caso de teste de alta prioridade encontrado."})
			return
		}
		w.WriteHeader(http.StatusOK)
		
		// Retorna os casos de teste de alta prioridade
		json.NewEncoder(w).Encode(higtPriorityTests)
	})
	// FIM DO TRECHO NOVO DO EPISÓDIO 6


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
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/globusdigital/soap"
)

func main() {
	// Persistência simples em JSON (arquivos locais).
	// Se não existirem, cria com estrutura vazia.
	ensureFile("pontos.json", []byte("{}\n"))

	// Migrações best-effort:
	// - users.json (antigo) -> clients.json (novo)
	// - pontos.json no formato {"pontos":[...]} -> {"<codigoEmbarcacao>":[...]}
	migratePontosIfNeeded("pontos.json")
	migrateUsersToClientsIfNeeded("users.json", "clients.json")

	// Se ainda não existir, cria vazio.
	ensureFile("clients.json", []byte("{}\n"))

	store := NewStore("clients.json", "pontos.json")

	s := soap.NewServer()
	// Desabilitar logging de requisições SOAP
	// s.Log = log.Println

	s.RegisterHandler(
		"/rastro",
		soapActionRastro,
		"registraPontos",
		func() any { return new(RegistraPontosRequest) },
		func(request any, w http.ResponseWriter, r *http.Request) (any, error) {
			req := request.(*RegistraPontosRequest)
			confirmacao, err := RegistraPontos(req.Autenticacao, req.CodigoEmbarcacao, req.PontosEmbarcacao.Items, store)
			if err != nil {
				return nil, err
			}
			return &RegistraPontosResponse{Confirmacao: confirmacao}, nil
		},
	)

	s.RegisterHandler(
		"/rastro",
		soapActionRastro,
		"consultaPontos",
		func() any { return new(ConsultaPontosRequest) },
		func(request any, w http.ResponseWriter, r *http.Request) (any, error) {
			req := request.(*ConsultaPontosRequest)
			pts, err := ConsultaPontos(req.Autenticacao, req.CodigoEmbarcacao, req.DataInicial.Time, req.DataFinal.Time, store)
			if err != nil {
				return nil, err
			}
			return &ConsultaPontosResponse{Pontos: Pontos{Items: pts}}, nil
		},
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("Rastro SOAP server em :%s (POST em /rastro)", port)
	log.Fatal(http.ListenAndServe(":"+port, s))
}

func ensureFile(path string, defaultContent []byte) {
	if _, err := os.Stat(path); err == nil {
		return
	}
	// best-effort: cria/reescreve se não existir
	_ = os.WriteFile(path, defaultContent, 0o644)
}

func migrateUsersToClientsIfNeeded(usersPath, clientsPath string) {
	// Se já existe e tem pelo menos 1 client, não mexe.
	if b, err := os.ReadFile(clientsPath); err == nil {
		var existing ClientsFile
		if err := json.Unmarshal(b, &existing); err == nil && len(existing) > 0 {
			return
		}
	}

	b, err := os.ReadFile(usersPath)
	if err != nil {
		return
	}
	// formato antigo:
	// {"users":[{"autenticacao":"...","embarcacoes":[...]}]}
	var legacy struct {
		Users []struct {
			Autenticacao string `json:"autenticacao"`
		} `json:"users"`
	}
	if err := json.Unmarshal(b, &legacy); err != nil {
		return
	}
	clients := ClientsFile{}
	for _, u := range legacy.Users {
		if u.Autenticacao == "" {
			continue
		}
		clients[u.Autenticacao] = Client{
			ID:          u.Autenticacao,
			Name:        "",
			Email:       "",
			Description: "",
		}
	}
	_ = atomicWriteJSON(clientsPath, clients)
}

func migratePontosIfNeeded(pontosPath string) {
	b, err := os.ReadFile(pontosPath)
	if err != nil || len(b) == 0 {
		return
	}

	// Detecta formato legado: {"pontos":[...]}
	var top map[string]json.RawMessage
	if err := json.Unmarshal(b, &top); err != nil {
		return
	}
	if raw, ok := top["pontos"]; !ok {
		// formato novo (map por embarcação) ou outro; não mexe.
		return
	} else {
		// precisa migrar; abaixo trata.
		_ = raw
	}

	// formato antigo:
	// {"pontos":[{"codigoEmbarcacao":"...","latitude":...,"longitude":...,"dataHora":"...","leituraSensores":[...]}]}
	var legacy struct {
		Pontos []struct {
			CodigoEmbarcacao string          `json:"codigoEmbarcacao"`
			Latitude         float64         `json:"latitude"`
			Longitude        float64         `json:"longitude"`
			DataHora         XSDDateTime     `json:"dataHora"`
			LeituraSensores  []LeituraSensor `json:"leituraSensores"`
		} `json:"pontos"`
	}
	if err := json.Unmarshal(b, &legacy); err != nil {
		return
	}

	out := PontosPorEmbarcacao{}
	for _, p := range legacy.Pontos {
		if p.CodigoEmbarcacao == "" {
			continue
		}
		out[p.CodigoEmbarcacao] = append(out[p.CodigoEmbarcacao], StoredPonto{
			Latitude:        p.Latitude,
			Longitude:       p.Longitude,
			DataHora:        p.DataHora,
			LeituraSensores: p.LeituraSensores,
		})
	}
	_ = atomicWriteJSON(pontosPath, out)
}

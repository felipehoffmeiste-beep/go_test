package main

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/globusdigital/soap"
)

// Client (clients.json)
// Formato desejado:
// {
//   "<uuid>": {"id":"<uuid>","name":"...","email":"...","description":"..."},
//   ...
// }
type Client struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

type ClientsFile map[string]Client

// PontosPorEmbarcacao (pontos.json)
// Formato desejado:
// {
//   "12345SC": [ {latitude, longitude, dataHora, leituraSensores:[...]} ],
//   ...
// }
type StoredPonto struct {
	Latitude        float64         `json:"latitude"`
	Longitude       float64         `json:"longitude"`
	DataHora        XSDDateTime     `json:"dataHora"`
	LeituraSensores []LeituraSensor `json:"leituraSensores,omitempty"`
}

type PontosPorEmbarcacao map[string][]StoredPonto

type Store struct {
	mu         sync.Mutex
	clientsPath string
	pontosPath string
}

func NewStore(clientsPath, pontosPath string) *Store {
	return &Store{clientsPath: clientsPath, pontosPath: pontosPath}
}

func (s *Store) loadClients() (ClientsFile, error) {
	b, err := os.ReadFile(s.clientsPath)
	if err != nil {
		return nil, err
	}
	if len(b) == 0 {
		return ClientsFile{}, nil
	}
	var cf ClientsFile
	if err := json.Unmarshal(b, &cf); err != nil {
		return nil, err
	}
	if cf == nil {
		cf = ClientsFile{}
	}
	return cf, nil
}

func (s *Store) loadPontos() (PontosPorEmbarcacao, error) {
	b, err := os.ReadFile(s.pontosPath)
	if err != nil {
		return nil, err
	}
	if len(b) == 0 {
		return PontosPorEmbarcacao{}, nil
	}
	var pf PontosPorEmbarcacao
	if err := json.Unmarshal(b, &pf); err != nil {
		return nil, err
	}
	if pf == nil {
		pf = PontosPorEmbarcacao{}
	}
	return pf, nil
}

func atomicWriteJSON(path string, v any) error {
	tmp := path + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		_ = f.Close()
		_ = os.Remove(tmp)
		return err
	}
	if err := f.Close(); err != nil {
		_ = os.Remove(tmp)
		return err
	}
	return os.Rename(tmp, path)
}

func (s *Store) authOK(autenticacao string) (bool, error) {
	cf, err := s.loadClients()
	if err != nil {
		return false, err
	}
	_, ok := cf[autenticacao]
	return ok, nil
}

func (s *Store) AppendPontos(autenticacao, codigoEmbarcacao string, pontos []Ponto) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	ok, err := s.authOK(autenticacao)
	if err != nil {
		return err
	}
	if !ok {
		return &soap.Fault{Code: "Client", String: "autenticação inválida"}
	}
	pf, err := s.loadPontos()
	if err != nil {
		return err
	}

	for _, p := range pontos {
		pf[codigoEmbarcacao] = append(pf[codigoEmbarcacao], StoredPonto{
			DataHora:   p.DataHora,
			Latitude:         p.Latitude,
			Longitude:        p.Longitude,
			LeituraSensores:  p.LeituraSensores.Items,
		})
	}
	return atomicWriteJSON(s.pontosPath, pf)
}

func (s *Store) QueryPontos(autenticacao, codigoEmbarcacao string, ini, fim time.Time) ([]Ponto, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ok, err := s.authOK(autenticacao)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, &soap.Fault{Code: "Client", String: "autenticação inválida"}
	}

	pf, err := s.loadPontos()
	if err != nil {
		return nil, err
	}

	stored := pf[codigoEmbarcacao]
	var out []Ponto
	for _, sp := range stored {
		dh := sp.DataHora.Time
		if dh.Before(ini) || dh.After(fim) {
			continue
		}
		out = append(out, Ponto{
			Latitude:  sp.Latitude,
			Longitude: sp.Longitude,
			DataHora:  sp.DataHora,
			LeituraSensores: LeituraSensores{
				Items: sp.LeituraSensores,
			},
		})
	}
	return out, nil
}

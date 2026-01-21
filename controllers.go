package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/globusdigital/soap"
)

// RegistraPontos implementa o método do ANEXO I:
// String registraPontos(String autenticacao, String codigoEmbarcacao, Ponto[] pontos)
func RegistraPontos(autenticacao, codigoEmbarcacao string, pontos []Ponto, store *Store) (string, error) {
	if autenticacao == "" || codigoEmbarcacao == "" {
		return "", &soap.Fault{Code: "Client", String: "autenticacao e codigoEmbarcacao são obrigatórios"}
	}
	if len(pontos) == 0 {
		return "", &soap.Fault{Code: "Client", String: "pontosEmbarcacao deve conter ao menos 1 Ponto"}
	}

	if err := store.AppendPontos(autenticacao, codigoEmbarcacao, pontos); err != nil {
		return "", err
	}

	// Máscara DDMMYYYYHHNNXXXXXXXXXX
	return makeConfirmacao(), nil
}

// ConsultaPontos implementa o método do ANEXO I:
// Ponto[] consultaPontos(String autenticacao, String codigoEmbarcacao, DateTime dataInicial, DateTime dataFinal)
// ... existing code ...

func ConsultaPontos(autenticacao, codigoEmbarcacao string, dataInicial, dataFinal time.Time, store *Store) ([]Ponto, error) {
	if autenticacao == "" || codigoEmbarcacao == "" {
		return nil, &soap.Fault{Code: "Client", String: "autenticacao e codigoEmbarcacao são obrigatórios"}
	}
	if dataFinal.Before(dataInicial) {
		return nil, &soap.Fault{Code: "Client", String: "dataFinal < dataInicial"}
	}

	// Restrição do documento: consultas limitadas às últimas 24h.
	if dataFinal.Sub(dataInicial) > 24*time.Hour {
		return nil, &soap.Fault{Code: "Client", String: "intervalo maior que 24h não permitido"}
	}

	now := time.Now().UTC()
	// Verificar se o intervalo completo está dentro das últimas 24 horas
	// dataInicial deve ser >= (now - 24h) e dataFinal deve ser <= now
	limite24h := now.Add(-24 * time.Hour)
	if dataInicial.Before(limite24h) {
		return nil, &soap.Fault{Code: "Client", String: "dataInicial fora da janela das últimas 24h"}
	}
	if dataFinal.After(now) {
		return nil, &soap.Fault{Code: "Client", String: "dataFinal não pode estar no futuro"}
	}

	return store.QueryPontos(autenticacao, codigoEmbarcacao, dataInicial, dataFinal)
}

// ... existing code ...

func makeConfirmacao() string {
	// DDMMYYYYHHNNXXXXXXXXXX
	now := time.Now().UTC()
	prefix := now.Format("020120061504")
	n := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1_000_000_0000)
	return fmt.Sprintf("%s%010d", prefix, n)
}

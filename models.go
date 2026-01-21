package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"time"
)

const soapActionRastro = "urn:RastroAction"

// --- dateTime (xsd:dateTime) ---
type XSDDateTime struct{ time.Time }

func (t *XSDDateTime) UnmarshalText(b []byte) error {
	s := string(b)
	// mais comum em SOAP: RFC3339 / RFC3339Nano
	if tt, err := time.Parse(time.RFC3339Nano, s); err == nil {
		t.Time = tt
		return nil
	}
	if tt, err := time.Parse(time.RFC3339, s); err == nil {
		t.Time = tt
		return nil
	}
	// fallback (sem timezone): assume UTC
	if tt, err := time.Parse("2006-01-02T15:04:05", s); err == nil {
		t.Time = tt.UTC()
		return nil
	}
	return fmt.Errorf("dateTime inválido: %q", s)
}

func (t XSDDateTime) MarshalText() ([]byte, error) {
	return []byte(t.UTC().Format(time.RFC3339)), nil
}

// --- modelos de domínio / SOAP ---
type LeituraSensor struct {
	TipoSensor uint32 `xml:"tipoSensor" json:"tipoSensor"`
	Valor      string `xml:"valor" json:"valor"`
}

type LeituraSensores struct {
	// No seu XML você envia `<LeituraSensor>...</LeituraSensor>`.
	// Aceitamos também `<item>` (alguns toolkits SOAP fazem isso).
	Items []LeituraSensor `xml:"LeituraSensor" json:"items,omitempty"`
}

func (ls *LeituraSensores) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var items []LeituraSensor
	for {
		tok, err := d.Token()
		if err == io.EOF {
			ls.Items = items
			return nil
		}
		if err != nil {
			return err
		}
		switch se := tok.(type) {
		case xml.StartElement:
			if se.Name.Local == "item" || se.Name.Local == "LeituraSensor" {
				var v LeituraSensor
				if err := d.DecodeElement(&v, &se); err != nil {
					return err
				}
				items = append(items, v)
			} else {
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if se.Name == start.Name {
				ls.Items = items
				return nil
			}
		}
	}
}

type Ponto struct {
	Latitude        float64         `xml:"latitude" json:"latitude"`
	Longitude       float64         `xml:"longitude" json:"longitude"`
	DataHora        XSDDateTime     `xml:"dataHora" json:"dataHora"`
	LeituraSensores LeituraSensores `xml:"leituraSensores" json:"leituraSensores,omitempty"`
}

type Pontos struct {
	// No seu XML você envia `<Ponto>...</Ponto>`.
	// Aceitamos também `<item>` (alguns toolkits SOAP fazem isso).
	Items []Ponto `xml:"Ponto" json:"items,omitempty"`
}

func (p *Pontos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var items []Ponto
	for {
		tok, err := d.Token()
		if err == io.EOF {
			p.Items = items
			return nil
		}
		if err != nil {
			return err
		}
		switch se := tok.(type) {
		case xml.StartElement:
			if se.Name.Local == "item" || se.Name.Local == "Ponto" {
				var v Ponto
				if err := d.DecodeElement(&v, &se); err != nil {
					return err
				}
				items = append(items, v)
			} else {
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if se.Name == start.Name {
				p.Items = items
				return nil
			}
		}
	}
}

// --- SOAP messages (RPC style) ---
type RegistraPontosRequest struct {
	XMLName          xml.Name `xml:"registraPontos"`
	Autenticacao     string   `xml:"autenticacao"`
	CodigoEmbarcacao string   `xml:"codigoEmbarcacao"`
	PontosEmbarcacao Pontos   `xml:"pontosEmbarcacao"`
}

type RegistraPontosResponse struct {
	XMLName     xml.Name `xml:"registraPontosResponse"`
	Confirmacao string   `xml:"confirmacao"`
}

type ConsultaPontosRequest struct {
	XMLName          xml.Name    `xml:"consultaPontos"`
	Autenticacao     string      `xml:"autenticacao"`
	CodigoEmbarcacao string      `xml:"codigoEmbarcacao"`
	DataInicial      XSDDateTime `xml:"dataInicial"`
	DataFinal        XSDDateTime `xml:"dataFinal"`
}

type ConsultaPontosResponse struct {
	XMLName xml.Name `xml:"consultaPontosResponse"`
	Pontos  Pontos   `xml:"pontos"`
}

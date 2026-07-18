package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func lerNoticias() ([]Noticia, error) {
	var noticias []Noticia

	dados, erro := os.ReadFile("noticias.json")

	if erro != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo: %w", erro)
	}

	erro = json.Unmarshal(dados, &noticias)

	if erro != nil {
		return nil, fmt.Errorf("erro ao converter o arquivo: %w", erro)
	}

	return noticias, nil
}

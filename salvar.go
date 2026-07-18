package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func salvarNoticias(noticias []Noticia) error {
	dados, erro := json.MarshalIndent(noticias, "", "  ")

	if erro != nil {
		return fmt.Errorf("Erro ao converter: %w", erro)
	}

	erro = os.WriteFile("noticias_salvas.json", dados, 0644)
	if erro != nil {
		return fmt.Errorf("Erro ao escrever os dados: %w", erro)
	}

	return nil
}

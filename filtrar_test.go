package main

import "testing"

func TestFiltrarNoticias(t *testing.T) {
	noticias := []Noticia{
		{
			Titulo:    "Claude ensina programação",
			Fonte:     "Anthropic",
			Categoria: "Tecnologia",
			Link:      "https://anthopic.com",
		},
		{
			Titulo:    "Golang, a linguagem do futuro",
			Fonte:     "GOlang",
			Categoria: "Tecnologia",
			Link:      "https://golang.com",
		},
		{
			Titulo:    "Flamengo é campeão de novo",
			Fonte:     "GE",
			Categoria: "Esporte",
			Link:      "https://globo.com",
		},
	}

	casos := []struct {
		nome              string
		categoria         string
		resultadoEsperado int
	}{
		{nome: "filtrar tecnologia", categoria: "Tecnologia", resultadoEsperado: 2},
		{nome: "filtrar esporte", categoria: "Esporte", resultadoEsperado: 1},
		{nome: "categoria inexistente", categoria: "Política", resultadoEsperado: 0},
		{nome: "filtrar culinária", categoria: "Culinária", resultadoEsperado: 0},
	}

	for _, caso := range casos {
		resultado := filtrarNoticias(noticias, caso.categoria)
		if len(resultado) != caso.resultadoEsperado {
			t.Errorf("%s: esperava %d, recebi %d", caso.nome, caso.resultadoEsperado, len(resultado))
		}
	}

}

package main

import "strings"

func filtrarNoticias(noticias []Noticia, categoria string) []Noticia {

	var resultado []Noticia

	for _, noticia := range noticias {

		if noticia.Categoria == categoria {
			resultado = append(resultado, noticia)
		}
	}

	return resultado

}

func filtrarPorPalavraChave(noticias []Noticia, palavraChave string) []Noticia {

	var resultado []Noticia

	for _, noticia := range noticias {

		if strings.Contains(strings.ToLower(noticia.Titulo), strings.ToLower(palavraChave)) {
			resultado = append(resultado, noticia)
		}
	}

	return resultado
}

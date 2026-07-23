package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

type RespostaAPI struct {
	Status string      `json:"status"`
	News   []ArtigoAPI `json:"news"`
}

type ArtigoAPI struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	Url         string   `json:"url"`
	Category    []string `json:"category"`
}

func buscarNoticiasAPI(categoria string) ([]Noticia, error) {
	chave := os.Getenv("CURRENTS_API_KEY")
	url := fmt.Sprintf("https://api.currentsapi.services/v2/latest-news?category=%s&apiKey=%s", categoria, chave)

	resposta, erro := http.Get(url)
	if erro != nil {
		return nil, fmt.Errorf("Erro na url: %w", erro)
	}
	defer resposta.Body.Close()

	dados, erro := io.ReadAll(resposta.Body)
	if erro != nil {
		return nil, fmt.Errorf("Erro ao ler o body: %w", erro)
	}
	var respostaAPI RespostaAPI

	erro = json.Unmarshal(dados, &respostaAPI)
	if erro != nil {
		return nil, fmt.Errorf("Erro ao converter JSON: %w", erro)
	}

	var noticias []Noticia
	for _, artigo := range respostaAPI.News {
		noticias = append(noticias, converterParaNoticia(artigo))
	}

	return noticias, nil
}

func converterParaNoticia(artigo ArtigoAPI) Noticia {
	categoria := "Sem categoria"

	if len(artigo.Category) > 0 {
		categoria = artigo.Category[0]
	}

	return Noticia{
		Titulo:    artigo.Title,
		Fonte:     artigo.Author,
		Categoria: categoria,
		Link:      artigo.Url,
	}

}

func buscarVariasCategorias(categorias []string) []Noticia {
	var wg sync.WaitGroup
	canal := make(chan []Noticia, len(categorias))
	for _, categoria := range categorias {
		wg.Add(1)
		go func() {
			defer wg.Done()
			noticias, erro := buscarNoticiasAPI(categoria)
			if erro == nil {
				canal <- noticias
			}
		}()
	}

	wg.Wait()
	close(canal)

	var todasNoticias []Noticia

	for noticias := range canal {
		todasNoticias = append(todasNoticias, noticias...)
	}

	return todasNoticias
}

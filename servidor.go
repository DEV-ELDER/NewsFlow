package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func criarPaginaNoticias(banco *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		noticias, erro := listarNoticiasSalvas(banco)
		if erro != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		categoria := r.URL.Query().Get("categoria")
		if categoria != "" {
			noticias = filtrarNoticias(noticias, categoria)
		}

		dadosJSON, erro := json.Marshal(noticias)
		if erro != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(dadosJSON)
	}

}

func criarPaginaAtualizar(banco *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		categorias := []string{"science_technology", "sport", "automotive", "politics_government"}
		noticias := buscarVariasCategorias(categorias)
		erro := salvarNoticiasBanco(noticias, banco)
		if erro != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Notícias atualizadas com sucesso")

	}

}

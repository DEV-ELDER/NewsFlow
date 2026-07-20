package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func conectarBanco() (*sql.DB, error) {
	stringConexao := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	banco, erro := sql.Open("postgres", stringConexao)

	if erro != nil {
		return nil, fmt.Errorf("erro de conexão com o banco: %w", erro)
	}

	erro = banco.Ping()
	if erro != nil {
		return nil, fmt.Errorf("Erro na tentativa de conexão com o banco:%w", erro)
	}

	return banco, nil

}

func salvarNoticiasBanco(noticias []Noticia, banco *sql.DB) error {

	for _, noticia := range noticias {
		_, erro := banco.Exec(
			"INSERT INTO noticias (titulo, fonte, categoria, link) VALUES ($1, $2, $3, $4) ON CONFLICT (link) DO NOTHING",
			noticia.Titulo, noticia.Fonte, noticia.Categoria, noticia.Link,
		)
		if erro != nil {
			return fmt.Errorf("Erro ao salvar noticia: %w", erro)

		}

	}

	return nil
}

func listarNoticiasSalvas(banco *sql.DB) ([]Noticia, error) {
	linhas, erro := banco.Query("SELECT titulo, fonte, categoria, link FROM noticias")
	if erro != nil {
		return nil, fmt.Errorf("Erro ao consultar notícias: %w", erro)
	}
	defer linhas.Close()
	var noticias []Noticia

	for linhas.Next() {
		var titulo, fonte,
			categoria, link string

		erro := linhas.Scan(&titulo, &fonte, &categoria, &link)
		if erro != nil {
			return nil, fmt.Errorf("Erro ao ler linha: %w", erro)
		}

		noticias = append(noticias, Noticia{
			Titulo:    titulo,
			Fonte:     fonte,
			Categoria: categoria,
			Link:      link,
		})
	}

	if erro := linhas.Err(); erro != nil {
		return nil, fmt.Errorf("Erro ao iterar linhas: %w", erro)
	}

	return noticias, nil
}

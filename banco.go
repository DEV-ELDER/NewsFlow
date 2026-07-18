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
			"INSERT INTO noticias (titulo, fonte, categoria, link) VALUES ($1, $2, $3, $4)",
			noticia.Titulo, noticia.Fonte, noticia.Categoria, noticia.Link,
		)
		if erro != nil {
			return fmt.Errorf("Erro ao salvar noticia: %w", erro)

		}

	}

	return nil
}

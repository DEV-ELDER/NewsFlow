package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {

	erro := godotenv.Load()
	if erro != nil {
		fmt.Println("Erro ao carregar .env:", erro)
	}

	banco, erro := conectarBanco()
	if erro != nil {
		fmt.Println("erro ao conectar no banco:", erro)
		return
	}
	fmt.Println("Conexão com o banco funcionando")

	noticias, erro := buscarNoticiasAPI()
	if erro != nil {
		fmt.Println("erro ao buscar notícias", erro)
		return
	}

	resultado := filtrarNoticias(noticias, "science_technology")
	listarNoticias(resultado)

	erro = salvarNoticiasBanco(resultado, banco)
	if erro != nil {
		fmt.Println("Erro ao salvar noticias", erro)
		return
	}

	noticiasSalvas, erro := listarNoticiasSalvas(banco)
	if erro != nil {
		fmt.Println("Erro ao listar notícias salvas:", erro)
		return
	}

	listarNoticias(noticiasSalvas)

}

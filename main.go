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

	categorias := []string{"science_technology", "sport", "automotive", "politics_government"}
	noticias := buscarVariasCategorias(categorias)

	//resultado := filtrarNoticias(noticias, "science_technology")
	listarNoticias(noticias)

	erro = salvarNoticiasBanco(noticias, banco)
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

package main

import "fmt"

func main() {

	noticias, erro := buscarNoticiasAPI()
	if erro != nil {
		fmt.Println("erro ao buscar notícias", erro)
		return
	}

	fmt.Println("Total de notícias recebidas:", len(noticias))
	for _, noticia := range noticias {
		fmt.Println("Categoria recebida:", noticia.Categoria)
	}

	resultado := filtrarNoticias(noticias, "science_technology")
	listarNoticias(resultado)

	erro = salvarNoticias(resultado)
	if erro != nil {
		fmt.Println("Erro ao salvar noticias", erro)
		return
	}

}

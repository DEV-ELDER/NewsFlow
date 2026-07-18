package main

import "fmt"

func listarNoticias(noticias []Noticia) {
	for _, noticia := range noticias {
		fmt.Println("========================")
		fmt.Println("Título:", noticia.Titulo)
		fmt.Println("Fonte:", noticia.Fonte)
		fmt.Println("Categoria:", noticia.Categoria)
		fmt.Println("Link:", noticia.Link)
		fmt.Println("========================")
	}
}

package main

func filtrarNoticias(noticias []Noticia, categoria string) []Noticia {

	var resultado []Noticia

	for _, noticia := range noticias {

		if noticia.Categoria == categoria {
			resultado = append(resultado, noticia)
		}
	}

	return resultado

}

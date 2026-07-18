package main

type Noticia struct {
	Titulo    string `json:"titulo"`
	Fonte     string `json:"fonte"`
	Categoria string `json:"categoria"`
	Link      string `json:"link"`
}

package main

import "testing"

type BancoFake struct {
	NoticiasSalvas []Noticia
}

func (b *BancoFake) Salvar(noticias []Noticia) error {

	b.NoticiasSalvas = append(b.NoticiasSalvas, noticias...)

	return nil
}

func TestBancoFakeSalvar(t *testing.T) {
	banco := BancoFake{}

	noticias := []Noticia{
		{Titulo: "Teste", Fonte: "Fonte Teste", Categoria: "Categoria Teste", Link: "https://teste.com"},
	}

	banco.Salvar(noticias)
	if len(banco.NoticiasSalvas) != 1 {
		t.Errorf("esperava 1 notícia salva, mas tinha %d", len(banco.NoticiasSalvas))
	}

}

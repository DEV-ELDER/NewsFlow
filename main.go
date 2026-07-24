package main

import (
	"fmt"
	"net/http"

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

	http.HandleFunc("/noticias", criarPaginaNoticias(banco))
	http.HandleFunc("/atualizar", criarPaginaAtualizar(banco))
	http.ListenAndServe(":8080", nil)

}

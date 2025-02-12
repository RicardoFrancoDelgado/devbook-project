package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	fmt.Println(config.StringConexao)
	r := router.Gerar()
	fmt.Println("Rodando API")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}

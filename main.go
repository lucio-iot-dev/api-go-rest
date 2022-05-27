package main

import (
	"api-go-rest/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}

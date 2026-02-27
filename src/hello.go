package main

import "fmt"

func main() {
	nome := "Miguel Silva"
	versao := 1.1
	fmt.Println("Olá, ", nome)

	fmt.Println("Menu Sistema \n\n 1 - Iniciar Programa \n 2 - Exibir Logs \n 3 - Sair \n\n")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("Endereço na memória para a var comando: ", &comando)

	fmt.Println("Opção selecionada: ", comando)

	fmt.Println("Versão do programa", versao)
}

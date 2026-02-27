package main

import (
	"fmt"
	"os"
)

func main() {
	systemInfoIn()
	menu()
	result(selectOption())
	systemInfoOut()
}

func menu() {
	fmt.Println("\nMenu Sistema")
	fmt.Println("1 - Iniciar Programa")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair")
}

func selectOption() int {
	var comando int
	fmt.Scan(&comando, "\n")
	fmt.Println("\nOpção selecionada: ", comando)
	return comando
}

func result(comando int) {
	mensagemOpcao1 := "Monitoramento iniciado."
	mensagemOpcao2 := "Exibir logs."
	mensagemOpcao3 := "Sair."
	mensagemOpcao4 := "Opção inválida."

	switch comando {
	case 1:
		fmt.Println(mensagemOpcao1)
	case 2:
		fmt.Println(mensagemOpcao2)
	case 3:
		fmt.Println(mensagemOpcao3)
		os.Exit(0)
	default:
		fmt.Println(mensagemOpcao4)
	}

	/**
	* por default, não é necessário o break
	* Se incluído seria apenas por didática
	 */
}

func systemInfoIn() {
	nome := "Miguel Silva"
	fmt.Println("Olá, ", nome)
}

func systemInfoOut() {
	versao := 1.1
	fmt.Println("\nVersão do programa", versao)
}

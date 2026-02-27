package main

import "fmt"

func main2() {
	nome := "Miguel Silva"
	fmt.Println("Olá, ", nome)

	fmt.Println("Menu Sistema \n\n 1 - Iniciar Monitoramentp \n 2 - Exibir Logs \n 3 - Sair \n\n")

	mensagemOpcao1 := "Monitoramento iniciado."
	mensagemOpcao2 := "Exibir logs."
	mensagemOpcao3 := "Sair."
	mensagemOpcao4 := "Opção inválida."
	var comando int

	fmt.Scan(&comando, "\n")
	fmt.Println("Opção selecionada: ", comando, "\n")

	/**
	* por default, não é necessário o break
	* Se incluído seria apenas por didática
	 */

	switch comando {
	case 1:
		fmt.Println(mensagemOpcao1)
	case 2:
		fmt.Println(mensagemOpcao2)
	case 3:
		fmt.Println(mensagemOpcao3)
	default:
		fmt.Println(mensagemOpcao4)
	}

	versao := 1.1
	fmt.Println("\nVersão do programa", versao)
}

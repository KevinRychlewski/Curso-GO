package main

import "fmt"

type usuario struct {
	nome string
	idade uint8 
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario 
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua aleatória",0}
	
	// SEGUNDO JEITO DE FAZER

	usuario2 := usuario{"Davi", 21,enderecoExemplo}
	fmt.Println(usuario2) 

	usuario3 := usuario{idade: 21}
	fmt.Println(usuario3)


}
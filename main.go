package main

import (
	"fmt"
	"banco/contas"
)

func main() {
	contaGuilherme := contas.ContaCorrente{Titular: "Guilherme", Saldo: 300}
	contaAriel := contas.ContaCorrente{Titular: "Ariel", Saldo: 300}

	status := contaGuilherme.Transferir(-200, &contaAriel)
	fmt.Println(status)
	fmt.Println(contaGuilherme)
	fmt.Println(contaAriel)
}

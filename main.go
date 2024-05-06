package main

import (
	//"banco/clientes"
	"banco/contas"
	"fmt"
)

func pagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaGuilherme := contas.ContaPoupanca{}
	contaGuilherme.Depositar(100)
	pagarBoleto(&contaGuilherme, 60)

	fmt.Println(contaGuilherme.ObterSaldo())

	contaAriel := contas.ContaCorrente{}
	contaAriel.Depositar(500)
	pagarBoleto(&contaAriel, 200)

	fmt.Println(contaAriel.ObterSaldo())

}

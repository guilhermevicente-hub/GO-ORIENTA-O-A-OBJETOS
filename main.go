package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valordoSaque float64) string {
	podeSacar := valordoSaque > 0 && valordoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valordoSaque
		return "Saque realizado com sucesso!!"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valordoDeposito float64) (string, float64) {
	if valordoDeposito > 0 {
		c.saldo += valordoDeposito
		return "Deposito realizado com sucesso!!", c.saldo
	} else {
		return "O valor do deposito menor que 0", c.saldo
	}

}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	contaGuilherme := ContaCorrente{titular: "Guilherme", saldo: 300}
	contaAriel := ContaCorrente{titular: "Ariel", saldo: 300}

	status := contaGuilherme.Transferir(-200, &contaAriel)
	fmt.Println(status)
	fmt.Println(contaGuilherme)
	fmt.Println(contaAriel)
}

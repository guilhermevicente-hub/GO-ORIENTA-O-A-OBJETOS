package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valordoSaque float64) string {
	podeSacar := valordoSaque > 0 && valordoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valordoSaque
		return "Saque realizado com sucesso!!"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valordoDeposito float64) (string, float64) {
	if valordoDeposito > 0 {
		c.saldo += valordoDeposito
		return "Deposito realizado com sucesso!!", c.saldo
	} else {
		return "O valor do deposito menor que 0", c.saldo
	}

}

func (c *ContaPoupanca) ObterSaldo() float64 {

	return c.saldo

}
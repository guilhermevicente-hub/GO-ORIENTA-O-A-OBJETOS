package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valordoSaque float64) string {
	podeSacar := valordoSaque > 0 && valordoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valordoSaque
		return "Saque realizado com sucesso!!"
	} else {
		return "saldo insuficiente"
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

func (c *ContaCorrente) ObterSaldo() float64 {

	return c.saldo

}

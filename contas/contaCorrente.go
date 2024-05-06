package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valordoSaque float64) string {
	podeSacar := valordoSaque > 0 && valordoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valordoSaque
		return "Saque realizado com sucesso!!"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valordoDeposito float64) (string, float64) {
	if valordoDeposito > 0 {
		c.Saldo += valordoDeposito
		return "Deposito realizado com sucesso!!", c.Saldo
	} else {
		return "O valor do deposito menor que 0", c.Saldo
	}

}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.Saldo && valorTransferencia > 0 {
		c.Saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

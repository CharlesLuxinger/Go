package contas

import "Go/ByteBank/clientes"

type ContaCorrente struct {
	Titular        clientes.Titular
	Agencia, Conta int
	saldo          float64
}

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}

func (c *ContaCorrente) Sacar(valor float64) (bool, float64) {
	if c.validaSaldoSuficiente(valor) {
		c.saldo -= valor
		return true, c.saldo
	}
	return false, c.saldo
}

func (c *ContaCorrente) Depositar(valor float64) (bool, float64) {
	if c.validaSaldoSuficiente(valor) {
		c.saldo += valor
		return true, c.saldo
	}
	return false, c.saldo
}

func (c *ContaCorrente) Transferir(contaDestino *ContaCorrente, valor float64) (bool, float64) {
	if c.validaSaldoSuficiente(valor) {
		c.Sacar(valor)
		contaDestino.Depositar(valor)
		return true, c.saldo
	}
	return false, c.saldo
}

func (c *ContaCorrente) validaSaldoSuficiente(valor float64) bool {
	return c.saldo >= valor && valor > 0
}

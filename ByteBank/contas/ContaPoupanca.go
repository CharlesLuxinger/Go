package contas

import "Go/ByteBank/clientes"

type ContaPoupanca struct {
	Titular        clientes.Titular
	Agencia, Conta int
	saldo          float64
}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}

func (c *ContaPoupanca) Sacar(valor float64) (bool, float64) {
	if c.validaSaldoSuficiente(valor) {
		c.saldo -= valor
		return true, c.saldo
	}
	return false, c.saldo
}

func (c *ContaPoupanca) Depositar(valor float64) (bool, float64) {
	if c.validaSaldoSuficiente(valor) {
		c.saldo += valor
		return true, c.saldo
	}
	return false, c.saldo
}

func (c *ContaPoupanca) Transferir(contaDestino *ContaPoupanca, valor float64) (bool, float64) {
	if c.validaSaldoSuficiente(valor) {
		c.Sacar(valor)
		contaDestino.Depositar(valor)
		return true, c.saldo
	}
	return false, c.saldo
}

func (c *ContaPoupanca) validaSaldoSuficiente(valor float64) bool {
	return c.saldo >= valor && valor > 0
}

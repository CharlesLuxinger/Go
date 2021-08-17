package main

import (
	"fmt"

	"Go/ByteBank/clientes"
	"Go/ByteBank/contas"
)

func main() {
	contaCharles := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:     "Charles",
			CPF:      "123.123.123-11",
			Profisso: "Programador",
		}}

	contaManuel := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:     "Manuel",
			CPF:      "321.123.123-11",
			Profisso: "Engenheiro",
		}}

	fmt.Println(contaCharles.GetSaldo())
	contaCharles.Depositar(500.)
	fmt.Println(contaCharles.GetSaldo())
	contaCharles.Sacar(200.)
	fmt.Println(contaCharles.GetSaldo())
	contaManuel.Depositar(500.)
	contaCharles.Transferir(&contaManuel, 200.)
	fmt.Println("Charles:", fmt.Sprintf("%10.2f", contaCharles.GetSaldo()), "Manuel:", fmt.Sprintf("%10.2f", contaManuel.GetSaldo()))

	PagarBoleto(&contaCharles, 150)
}

type Conta interface {
	Sacar(valor float64) (bool, float64)
}

func PagarBoleto(conta Conta, value float64) {
	conta.Sacar(value)
}
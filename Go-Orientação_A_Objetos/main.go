package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	contaDoWillyan := ContaCorrente{titular: "Willyan",
		numeroAgencia: 0001, numeroConta: 155478, saldo: 135.5}

	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	fmt.Println(contaDoWillyan)
	fmt.Println(contaDaBruna)

	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Sacar(400))
	fmt.Println(contaDaSilvia.saldo)
}

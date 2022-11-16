package main

import "fmt"

type DadosConta struct {
	Titular string  `json:"Titular"`
	Agencia int     `json:"Agência"`
	Conta   int     `json:"Conta"`
	Saldo   float64 `json:"Saldo"`
}

func (d *DadosConta) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= d.Saldo

	if podeSacar {
		resultado := d.Saldo - valorDoSaque
		fmt.Printf("Saque realizado com sucesso no valor de R$%v Reais, Saldo restante R$ %v Reais\n", valorDoSaque, resultado)
	} else {
		fmt.Println("Saldo insuficiente.")
	}
	return ""
}

func main() {

	PrimeiroCliente := DadosConta{"Guilherme Dias", 589, 123456, 12500}

	SegundoCliente := DadosConta{
		Titular: "Luciene Silva",
		Agencia: 589,
		Conta:   111233,
		Saldo:   489.90,
	}

	TerceiroCliente := DadosConta{
		Titular: "Diego Chaves",
		Agencia: 589,
		Conta:   543678,
		Saldo:   9890.76,
	}

	QuartoCliente := DadosConta{"Bruna Silva", 289, 111223, 87672.90}

	fmt.Println(PrimeiroCliente)
	fmt.Println(SegundoCliente)
	fmt.Println(TerceiroCliente)
	fmt.Println(QuartoCliente)

	contaDasilvia := DadosConta{}
	contaDasilvia.Titular = "Silvia"
	contaDasilvia.Saldo = 500.

	contaDasilvia.Sacar(-600)
	PrimeiroCliente.Sacar(200)
	SegundoCliente.Sacar(500)

}

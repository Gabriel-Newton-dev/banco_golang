package main

import (
	"fmt"

	"github.com/Gabriel-Newton-dev/banco_golang/contas"
)

func main() {

	PrimeiroCliente := contas.DadosConta{
		Titular: "Guilherme Dias",
		Agencia: 589,
		Conta:   123456,
		Saldo:   12500,
	}

	SegundoCliente := contas.DadosConta{
		Titular: "Luciene Silva",
		Agencia: 589,
		Conta:   111233,
		Saldo:   489.90,
	}

	TerceiroCliente := contas.DadosConta{
		Titular: "Diego Chaves",
		Agencia: 589,
		Conta:   543678,
		Saldo:   9890.76,
	}

	QuartoCliente := contas.DadosConta{
		Titular: "Bruna Silva",
		Agencia: 289,
		Conta:   111223,
		Saldo:   87672.90,
	}

	fmt.Println(PrimeiroCliente)
	fmt.Println(SegundoCliente)
	fmt.Println(TerceiroCliente)
	fmt.Println(QuartoCliente)

	contaDasilvia := contas.DadosConta{}
	contaDasilvia.Titular = "Silvia"
	contaDasilvia.Saldo = 500

	fmt.Println(contaDasilvia.Deposito(1000))
	//fmt.Println(contaDasilvia.Sacar(100))

	contaDasilvia.Tranferencia(1000, &PrimeiroCliente)

	fmt.Printf("O Saldo atualizado de %v é de R$ %v Reais\n", contaDasilvia.Titular, contaDasilvia.Saldo)
	fmt.Printf("O Saldo atualizado de %s é de R$ %v Reais\n", PrimeiroCliente.Titular, PrimeiroCliente.Saldo)

}

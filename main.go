package main

import (
	"fmt"

	"github.com/Gabriel-Newton-dev/banco_golang/clientes"
	"github.com/Gabriel-Newton-dev/banco_golang/contas"
)

func main() {

	DadosCliente.PrimeiroCliente

	SegundoCliente := contas.DadosConta{
		Titular: clientes.Titular{"Luciene Silva", "345.678.907-56", "Do lar"},
		Agencia: 589,
		Conta:   111233,
		Saldo:   489.90,
	}

	TerceiroCliente := contas.DadosConta{
		Titular: clientes.Titular{"Diego Chaves", "234.123.432-09", "Vendedor"},
		Agencia: 589,
		Conta:   543678,
		Saldo:   9890.76,
	}

	QuartoCliente := contas.DadosConta{
		Titular: clientes.Titular{"Bruna Silva", "259.696.696-56", "Secretária"},
		Agencia: 289,
		Conta:   111223,
		Saldo:   87672.90,
	}

	fmt.Println(SegundoCliente)
	fmt.Println(TerceiroCliente)
	fmt.Println(QuartoCliente)

	contaDasilvia := contas.DadosConta{}
	contaDasilvia.Titular = clientes.Titular{"Silvia", "345.383.898-90", "Desenvolvedora"}
	contaDasilvia.Saldo = 500

	fmt.Println(contaDasilvia.Deposito(1000))
	//fmt.Println(contaDasilvia.Sacar(100))

	contaDasilvia.Tranferencia(1000, &PrimeiroCliente)

	fmt.Printf("O Saldo atualizado de %v é de R$ %v Reais\n", contaDasilvia.Titular, contaDasilvia.Saldo)
	fmt.Printf("O Saldo atualizado de %s é de R$ %v Reais\n", PrimeiroCliente.Titular, PrimeiroCliente.Saldo)

}

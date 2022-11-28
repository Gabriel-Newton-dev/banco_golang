package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Gabriel-Newton-dev/banco_golang/clientes"
	"github.com/Gabriel-Newton-dev/banco_golang/contas"
)

func main() {

	PrimeiroCliente := contas.DadosConta{
		Titular: clientes.Titular{Nome: "Guilherme Dias", CPF: "123.456.789-00", Profissao: "Médico"},
		Agencia: 589,
		Conta:   123456,
		Saldo:   12500,
	}

	SegundoCliente := contas.DadosConta{
		Titular: clientes.Titular{Nome: "Luciene Silva", CPF: "345.678.907-56", Profissao: "Do lar"},
		Agencia: 589,
		Conta:   111233,
		Saldo:   489.90,
	}

	TerceiroCliente := contas.DadosConta{
		Titular: clientes.Titular{Nome: "Diego Chaves", CPF: "234.123.432-09", Profissao: "Vendedor"},
		Agencia: 589,
		Conta:   543678,
		Saldo:   9890.76,
	}

	QuartoCliente := contas.DadosConta{
		Titular: clientes.Titular{Nome: "Bruna Silva", CPF: "259.696.696-56", Profissao: "Secretária"},
		Agencia: 289,
		Conta:   111223,
		Saldo:   87672.90,
	}

	primeiroClienteJson, err := json.Marshal(PrimeiroCliente)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(primeiroClienteJson))
	fmt.Println(SegundoCliente)
	fmt.Println(TerceiroCliente)
	fmt.Println(QuartoCliente)

	contaDasilvia := contas.DadosConta{}
	contaDasilvia.Titular = clientes.Titular{Nome: "Silvia", CPF: "345.383.898-90", Profissao: "Desenvolvedora"}
	contaDasilvia.Saldo = 500

	fmt.Println(contaDasilvia.Deposito(1000))
	//fmt.Println(contaDasilvia.Sacar(100))

	contaDasilvia.Tranferencia(1000, &PrimeiroCliente)

	fmt.Printf("O Saldo atualizado de %v é de R$ %v Reais\n", contaDasilvia.Titular, contaDasilvia.Saldo)
	fmt.Printf("O Saldo atualizado de %s é de R$ %v Reais\n", PrimeiroCliente.Titular, PrimeiroCliente.Saldo)

}

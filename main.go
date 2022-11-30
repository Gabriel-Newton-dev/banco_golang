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

	contaDasilvia.Tranferencia(1000, &PrimeiroCliente)

	contaDasilvia.ObterSaldo()
	PrimeiroCliente.ObterSaldo()

	ContaDoDenis := contas.ContaPoupanca{
		Titular: clientes.Titular{
			Nome:      "Carlos Denis",
			CPF:       "432-543-323-76",
			Profissao: "Técnico de Informática",
		},
		NumeroAgencia: 256,
		NumeroConta:   7860,
		Operacao:      1,
		Saldo:         356,
	}

	ContaDoDenis.Depositar(100)
	ContaDoDenis.Sacar(300)
	ContaDoDenis.ValorDoSaldo()
}

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

func (d *DadosConta) Deposito(valorDoDeposito float64) (string, float64, string) {
	if valorDoDeposito > 0 {
		d.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso R$", d.Saldo, "reais"
	} else {
		return "Não é possível fazer deposito com valor negativo, seu saldo atual é de R$", d.Saldo, "Reais"
	}

}

func (d *DadosConta) tranferencia(valorDaTransferencia float64, contaDestino *DadosConta) bool {
	if valorDaTransferencia < d.Saldo {
		d.Saldo -= valorDaTransferencia
		contaDestino.Deposito((valorDaTransferencia))
		return true
	} else {
		return false
	}
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
	contaDasilvia.Saldo = 500

	fmt.Println(contaDasilvia.Deposito(1000))
	//fmt.Println(contaDasilvia.Sacar(100))

	tranferir := contaDasilvia.tranferencia(1000, &PrimeiroCliente)
	fmt.Println(tranferir)

	fmt.Printf("O Saldo atualizado de %v é de R$ %v Reais\n", contaDasilvia.Titular, contaDasilvia.Saldo)
	fmt.Printf("O Saldo atualizado de %s é de R$ %v Reais\n", PrimeiroCliente.Titular, PrimeiroCliente.Saldo)

}

package contas

import (
	"fmt"

	"github.com/Gabriel-Newton-dev/banco_golang/clientes"
)

type ContaPoupanca struct {
	Titular       clientes.Titular `json:"Titular"`
	NumeroAgencia int              `json:"Agência"`
	NumeroConta   int              `json:"Conta"`
	Operacao      int              `json:"Operação"`
	Saldo         float64          `json:"Saldo"`
}

// vou criar 2 possíveis transações bancárias, Sacar e depoistar

func (c *ContaPoupanca) Depositar(ValorParaDeposito float64) {
	podeDepositar := ValorParaDeposito > 0
	if podeDepositar {
		c.Saldo += ValorParaDeposito
		fmt.Printf("Sr(a) %s o deposito de R$ %v Reais foi realizado com sucesso, seu saldo atual é de R$ %v Reais.\n", c.Titular.Nome, ValorParaDeposito, c.Saldo)
	} else {
		fmt.Println("Deposito não realizado, favor verificar os dados")
	}
}

func (c *ContaPoupanca) Sacar(ValorDoSaque float64) {
	podeSacar := c.Saldo >= ValorDoSaque && ValorDoSaque > 0
	if podeSacar {
		c.Saldo -= ValorDoSaque
		fmt.Printf("Sr(a) %s seu saques no valor de R$ %v Reais foi realizado com sucesso, seu saldo atual agora é de R$ %v Reais.\n", c.Titular.Nome, ValorDoSaque, c.Saldo)
	} else {
		fmt.Println("Saque não realizado.")
	}
}

func (c *ContaPoupanca) ValorDoSaldo() {
	fmt.Printf("Sr(a) %s seu saldo atual é de R$ %v Reias.\n", c.Titular.Nome, c.Saldo)

}

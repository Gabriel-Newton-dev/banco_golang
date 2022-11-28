package contas

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

func (d *DadosConta) Tranferencia(valorDaTransferencia float64, contaDestino *DadosConta) bool {
	if valorDaTransferencia < d.Saldo {
		d.Saldo -= valorDaTransferencia
		contaDestino.Deposito((valorDaTransferencia))
		return true
	} else {
		return false
	}
}

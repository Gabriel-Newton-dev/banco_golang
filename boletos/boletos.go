package boletos

func PagarBoleto(conta VerficarConta, valorDoboleto float64) {
	conta.Sacar(valorDoboleto)
}

type VerficarConta interface {
	Sacar(valor float64)
}

package dadosclientes

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Gabriel-Newton-dev/banco_golang/clientes"
	"github.com/Gabriel-Newton-dev/banco_golang/contas"
)

func DadosCliente() {

	PrimeiroCliente := contas.DadosConta{
		Titular: clientes.Titular{"Guilherme Dias", "123.456.789-00", "Médico"},
		Agencia: 589,
		Conta:   123456,
		Saldo:   12500,
	}

	primeiroClienteJson, err := json.Marshal(PrimeiroCliente)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(primeiroClienteJson))

}

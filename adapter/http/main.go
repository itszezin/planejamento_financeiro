package http

import (
	"fmt"
	"net/http"

	"github.com/itszezin/planejamento_financeiro/adapter/http/actuator"
	"github.com/itszezin/planejamento_financeiro/adapter/http/transaction"
)

func Init() {

	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)
	http.HandleFunc("/health", actuator.Health)
	http.ListenAndServe(":8080", nil)

	fmt.Printf("Servidor iniciado em http://localhost:8080")
}

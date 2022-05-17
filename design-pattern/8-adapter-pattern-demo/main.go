package main

import (
	"fmt"

	"lecture/design-pattern/8-adapter-pattern-demo/payment"
)

func main() {
	clientRequest := payment.PaymentClientRequest{
		ClientID: "client-id",
		Amount:   10000,
	}

	clientPaymentAdapter := payment.PaymentAdapterLegacy{}

	amount, returnStr := clientPaymentAdapter.Pay(clientRequest)

	fmt.Println("Amount:", amount)
	fmt.Println("Status:", returnStr)

	fmt.Println("Refund:", clientPaymentAdapter.Refund(clientRequest))
}

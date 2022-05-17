package main

import (
	"fmt"

	"lecture/design-pattern/8-adapter-pattern/payment"
)

func main() {
	clientPaymentRequest := payment.ClientPaymentRequest{
		ClientID: "client-id",
		Amount:   100,
	}

	clientPaymentAdapter := payment.PaymentAdapter3rdParty{}

	amount, s := clientPaymentAdapter.Pay(clientPaymentRequest)

	fmt.Println("Amount:", amount)
	fmt.Println("Status:", s)
}

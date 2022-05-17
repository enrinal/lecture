package payment

// create adapter pattern

type ClientPaymentRequest struct {
	ClientID string
	Amount   int
}

type ClientPayment interface {
	Pay(request ClientPaymentRequest) (int, string)
}

type PaymentAdapter3rdParty struct {
	PaypalPayment PaypalPayment
}

func (p *PaymentAdapter3rdParty) Pay(request ClientPaymentRequest) (int, string) {
	paypalRequest := PaypalPaymentRequest{
		Amount:   float64(request.Amount),
		ClientID: request.ClientID,
	}
	amount, s := p.PaypalPayment.PayWithPaypal(paypalRequest)

	return int(amount), s
}

type PaypalPayment struct{}

type PaypalPaymentRequest struct {
	ClientID string
	Amount   float64
}

func (p *PaypalPayment) PayWithPaypal(request PaypalPaymentRequest) (float64, string) {
	return request.Amount, "client with id " + request.ClientID + " paid with paypal"
}

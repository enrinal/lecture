package payment

type PaymentClientRequest struct {
	ClientID string
	Amount   int
}

type PaymentClient interface {
	Pay(request PaymentClientRequest) (int, string)
	Refund(request PaymentClientRequest) string
}

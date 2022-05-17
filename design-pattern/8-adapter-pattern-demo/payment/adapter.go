package payment

type PaymentAdapterLegacy struct {
	PaymentLegacy PaymentLegacy
}

func (p *PaymentAdapterLegacy) Pay(request PaymentClientRequest) (int, string) {
	payWithLegacyReq := LegacyPaymentRequest{
		ClientID: request.ClientID,
		Amount:   float64(request.Amount),
	}

	amount, returnStr := p.PaymentLegacy.PayWithLegacy(payWithLegacyReq)

	return int(amount), returnStr
}

func (p *PaymentAdapterLegacy) Refund(request PaymentClientRequest) string {
	payWithLegacyReq := LegacyPaymentRequest{
		ClientID: request.ClientID,
		Amount:   float64(request.Amount),
	}

	return p.PaymentLegacy.RefundWithLegacy(payWithLegacyReq)
}

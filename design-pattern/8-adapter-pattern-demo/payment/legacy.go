package payment

import "fmt"

type PaymentLegacy struct{}

type LegacyPaymentRequest struct {
	ClientID string
	Amount   float64
}

func (pl *PaymentLegacy) PayWithLegacy(request LegacyPaymentRequest) (float64, string) {
	return request.Amount, "client with id " + request.ClientID + " paid with legacy"
}

func (pl *PaymentLegacy) RefundWithLegacy(request LegacyPaymentRequest) string {
	amountStr := fmt.Sprintf("%.2f", request.Amount)

	return "client with id " + request.ClientID + " refund with amount " + amountStr + " with legacy"
}

package payment

import stripe "github.com/stripe/stripe-go/v76"

type PaymentService struct {
}

type PaymenterService interface {
	Checkout(email string) (*stripe.CheckoutSession, error)
}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

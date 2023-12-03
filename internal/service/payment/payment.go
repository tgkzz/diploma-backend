package payment

import (
	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/customer"
)

var PriceId = "price_1OIkgVB9lxIKH7nckOfM9JZ8"

func (s *PaymentService) Checkout(email string) (*stripe.CheckoutSession, error) {
	// var discounts []*stripe.CheckoutSessionDiscountParams

	// discounts = []*stripe.CheckoutSessionDiscountParams{
	// 	&stripe.CheckoutSessionDiscountParams{
	// 		Coupon: stripe.String("FMARC"),
	// 	},
	// }

	customerParams := &stripe.CustomerParams{
		Email: stripe.String(email),
	}
	customerParams.AddMetadata("FinalEmail", email)
	newCustomer, err := customer.New(customerParams)

	if err != nil {
		return nil, err
	}

	meta := map[string]string{
		"FinalEmail": email,
	}

	params := &stripe.CheckoutSessionParams{
		Customer:   &newCustomer.ID,
		SuccessURL: stripe.String("https://www.youtube.com/channel/UCzgn3FvGR1UK_0M0B6GiLug"),
		CancelURL:  stripe.String("https://www.youtube.com/channel/UCzgn3FvGR1UK_0M0B6GiLug"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Mode: stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				Price:    stripe.String(PriceId),
				Quantity: stripe.Int64(1),
			},
		},
		SubscriptionData: &stripe.CheckoutSessionSubscriptionDataParams{
			TrialPeriodDays: stripe.Int64(7),
			Metadata:        meta,
		},
	}
	return session.New(params)
}

package internal

import "calcFeeSample/internal/app"

type Internal struct {
	CalcFee       app.CalcFee       // has-aなので委譲
	PaymentReport app.PaymentReport // has-aなので委譲
}

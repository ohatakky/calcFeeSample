package internal

import "calcFeeSample/internal/interface/api"

type Internal struct {
	CalcFee       api.CalcFee       // has-aなので委譲
	PaymentReport api.PaymentReport // has-aなので委譲
}

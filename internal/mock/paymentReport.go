package mock

import (
	"calcFeeSample/internal/interface/api"
)

type PaymentReport struct {
	api.PaymentReport // モックはis-aなので埋め込み
}

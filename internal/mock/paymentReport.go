package mock

import "calcFeeSample/internal/app"

type PaymentReport struct {
	app.PaymentReport // モックはis-aなので埋め込み
}

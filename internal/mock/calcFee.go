package mock

import "calcFeeSample/internal/app"

type CalcFee struct {
	app.CalcFee // モックはis-aなので埋め込み
}

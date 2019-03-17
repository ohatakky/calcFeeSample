package mock

import "calcFeeSample/external/app"

type CalcFee struct {
	app.CalcFee // モックはis-aなので埋め込み
}

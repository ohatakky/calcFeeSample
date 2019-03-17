package mock

import "calcFeeSample/internal/interface/api"

type CalcFee struct {
	api.CalcFee // モックはis-aなので埋め込み
}

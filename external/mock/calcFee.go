package mock

import "calcFeeSample/external/interface/api"

type CalcFee struct {
	api.CalcFee // モックはis-aなので埋め込み
}

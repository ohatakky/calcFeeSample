package main

import (
	"calcFeeSample/external"
	ex_mock "calcFeeSample/external/mock"
	"calcFeeSample/internal"
	"calcFeeSample/internal/mock"
	in_mock "calcFeeSample/internal/mock"
	"fmt"
)

func main() {
	ex_mock := ex_mock.CalcFee{}
	ex := external.External{CalcFee: ex_mock}
	fmt.Println(ex)

	in_mock_cf := in_mock.CalcFee{}
	in_mock_pr := mock.PaymentReport{}
	in := internal.Internal{CalcFee: in_mock_cf, PaymentReport: in_mock_pr}
	fmt.Println(in)
}

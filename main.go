package main

import (
	"calcFeeSample/external"
	"calcFeeSample/internal"
	"calcFeeSample/internal/mock"
	"fmt"
)

func main() {
	ex_mock := mock.CalcFee{}
	ex := external.External{CalcFee: ex_mock}
	fmt.Println(ex)

	in_mock_cf := mock.CalcFee{}
	in_mock_pr := mock.PaymentReport{}
	in := internal.Internal{CalcFee: in_mock_cf, PaymentReport: in_mock_pr}
	fmt.Println(in)
}

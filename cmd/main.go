/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	promptpayqr "github.com/kazekim/promptpay-qr-go"
)

func main() {

	qr := promptpayqr.New()
	amount := "420"
	payload := qr.GeneratePayload("0899999999", &amount)
	fmt.Println(payload)
}

/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package promptpayqr

import "github.com/skip2/go-qrcode"

func QRForTargetWithAmount(target, amount string) (*[]byte, error) {

	qr := New()
	payload := qr.GeneratePayload(target, &amount)

	var png []byte
	png, err := qrcode.Encode(payload, qrcode.Medium, 256)

	if err != nil {
		return nil, err
	}
	return &png, nil
}


func QRForTarget(target, amount string) (*[]byte, error) {

	qr := New()
	payload := qr.GeneratePayload(target, nil)

	var png []byte
	png, err := qrcode.Encode(payload, qrcode.Medium, 256)


	if err != nil {
		return nil, err
	}
	return &png, nil
}
/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package promptpayqr

import (
	"bytes"
	"fmt"
	"github.com/skip2/go-qrcode"
	"image"
	"image/png"
	"os"
)

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

func QRForBillPayment(billerID string, ref1 string, ref2 string, terminalID string, amount string) (*[]byte, error) {

	qr := New()
	payload := qr.GenerateBillPaymentPayload(billerID, ref1, ref2, &terminalID, &amount)

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

func ByteToImagePNG(imgByte []byte, filename string) error {
	img, _, _ := image.Decode(bytes.NewReader(imgByte))
	//save the imgByte to file
	out, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = png.Encode(out, img)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

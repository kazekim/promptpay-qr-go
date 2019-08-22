/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package promptpayqr

import (
	"flag"
	"fmt"
	"github.com/divan/qrlogo"
	"github.com/skip2/go-qrcode"
	"image"
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

func QRForTarget(target string) (*[]byte, error) {

	qr := New()
	payload := qr.GeneratePayload(target, nil)

	var png []byte
	png, err := qrcode.Encode(payload, qrcode.Medium, 256)

	if err != nil {
		return nil, err
	}
	return &png, nil
}


func QRWithPromptpayLogoForTargetWithAmount(target, amount string) (*[]byte, error) {

	var (
		input  = flag.String("promptpay", "promptpay.png", "Prompt Pay Logo")
		size   = flag.Int("size", 256, "Image size in pixels")
	)
	qr := New()
	payload := qr.GeneratePayload(target, &amount)

	file, err := os.Open(*input)
	if err != nil {
		fmt.Println("Failed to open logo:", err)
		os.Exit(1)
	}
	defer file.Close()

	logo, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Failed to decode PNG with logo:", err)
		os.Exit(1)
	}

	qrImage, err := qrlogo.Encode(payload, logo, *size)
	if err != nil {
		fmt.Println("Failed to encode QR:", err)
		os.Exit(1)
	}

	qrBytes := qrImage.Bytes()
	return &qrBytes, err
}

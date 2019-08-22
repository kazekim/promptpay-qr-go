package promptpayqr_test

import (
	"github.com/disintegration/imaging"
	"github.com/divan/qrlogo"
	tqrc "github.com/kazekim/promptpay-qr-go"
	"github.com/stretchr/testify/assert"
	"image"
	"testing"
)

func TestGenerateBillPaymentPayloadMustValid(t *testing.T) {
	billerID := "123456789012345"
	ref1 := "REF001"
	ref2 := "REF2"
	terminalID := "ABC001"
	amount := "555.55"
	expectedPayload := "00020101021230570016A00000067701011201151234567890123450206REF0010304REF253037645406555.555802TH62100706ABC0016304D3C4"
	qr := tqrc.New()
	actualPayload := qr.GenerateBillPaymentPayload(billerID, ref1, ref2, &terminalID, &amount)
	assert.Equal(t, expectedPayload, actualPayload)
}

func TestGenerateBillPaymentQRWithLogoMustHaveNoError(t *testing.T) {
	billerID := "123456789012345"
	ref1 := "REF001"
	ref2 := "REF2"
	terminalID := "ABC001"
	amount := "555.55"
	const QRImgFilename = "QRImg.png"
	const QRImgWithLogoFilename = "QRImgWithLogo.png"

	qr := tqrc.New()
	payload := qr.GenerateBillPaymentPayload(billerID, ref1, ref2, &terminalID, &amount)
	logo, err := imaging.Open("tqrp_logo.png")
	assert.Nil(t, err)
	png, err := qrlogo.Encode(payload, logo, 512)
	tqrc.ByteToImagePNG(png.Bytes(), QRImgFilename)
	src, err := imaging.Open(QRImgFilename)
	assert.Nil(t, err)
	src = imaging.Overlay(src, logo, image.Point{src.Bounds().Dx()/2 - (logo.Bounds().Dx() / 2), src.Bounds().Dy()/2 - (logo.Bounds().Dy()/2 + 8)}, 1.0)
	err = imaging.Save(src, QRImgWithLogoFilename)
	assert.Nil(t, err)
}

package promptpayqr_test

import (
	tqrc "github.com/kazekim/promptpay-qr-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateBillPaymentPayloadMustValid(t *testing.T) {
	billerID := "311040039475101" // SCB
	ref1 := "REF001"
	ref2 := "REF2"
	terminalID := "SCB001"
	amount := "555.55"
	expectedPayload := "00020101021230570016A00000067701011201153110400394751010206REF0010304REF253037645406555.555802TH62100706SCB001630437C6" // SCB
	qr := tqrc.New()
	actualPayload := qr.GenerateBillPaymentPayload(billerID, ref1, ref2, &terminalID, &amount)
	assert.Equal(t, expectedPayload, actualPayload)
}

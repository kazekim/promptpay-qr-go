/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package promptpayqr

import (
	"fmt"
	"github.com/kazekim/promptpay-qr-go/crc16"
	"regexp"
	"strconv"
	"strings"
)

const (
	ID_PAYLOAD_FORMAT                        = "00"
	ID_POI_METHOD                            = "01"
	ID_MERCHANT_INFORMATION_BOT              = "29"
	ID_MERCHANT_INFORMATION_BOT_BILL_PAYMENT = "30"
	ID_TRANSACTION_CURRENCY                  = "53"
	ID_TRANSACTION_AMOUNT                    = "54"
	ID_COUNTRY_CODE                          = "58"
	ID_DATA_OBJECTS                          = "62"
	ID_CRC                                   = "63"

	PAYLOAD_FORMAT_EMV_QRCPS_MERCHANT_PRESENTED_MODE = "01"
	POI_METHOD_STATIC                                = "11"
	POI_METHOD_DYNAMIC                               = "12"
	MERCHANT_INFORMATION_TEMPLATE_ID_GUID            = "00"
	BOT_ID_MERCHANT_PHONE_NUMBER                     = "01"
	BOT_ID_MERCHANT_TAX_ID                           = "02"
	BOT_ID_MERCHANT_EWALLET_ID                       = "03"
	BOT_ID_TAG30_AID                                 = "00"
	BOT_ID_TAG30_BILLER_ID                           = "01"
	BOT_ID_TAG30_REF1                                = "02"
	BOT_ID_TAG30_REF2                                = "03"
	BOT_ID_TAG62_TERMINAL_ID                         = "07"
	GUID_PROMPTPAY                                   = "A000000677010111"
	GUID_PROMPTPAY_BILL_PAYMENT                      = "A000000677010112"
	TRANSACTION_CURRENCY_THB                         = "764"
	COUNTRY_CODE_TH                                  = "TH"
)

type PromptPayQR struct {
}

func New() *PromptPayQR {
	return &PromptPayQR{}
}

func (qr *PromptPayQR) GeneratePayload(target string, amount *string) string {

	target = sanitizeTarget(target)

	var targetType string
	switch length := len(target); {
	case length >= 15:
		targetType = BOT_ID_MERCHANT_EWALLET_ID
	case length >= 13:
		targetType = BOT_ID_MERCHANT_TAX_ID
	default:
		targetType = BOT_ID_MERCHANT_PHONE_NUMBER
	}

	var data []string
	data = append(data,
		f(ID_PAYLOAD_FORMAT, PAYLOAD_FORMAT_EMV_QRCPS_MERCHANT_PRESENTED_MODE),
		f(ID_POI_METHOD, ifThenElse(amount != nil, POI_METHOD_DYNAMIC, POI_METHOD_STATIC).(string)),
		f(ID_MERCHANT_INFORMATION_BOT, serialize(
			[]string{
				f(MERCHANT_INFORMATION_TEMPLATE_ID_GUID, GUID_PROMPTPAY),
				f(targetType, formatTarget(target)),
			})),

		f(ID_COUNTRY_CODE, COUNTRY_CODE_TH),
		f(ID_TRANSACTION_CURRENCY, TRANSACTION_CURRENCY_THB),
	)

	if amount != nil {
		data = append(data, f(ID_TRANSACTION_AMOUNT, formatAmount(*amount)))
	}

	dataToCrc := serialize(data) + ID_CRC + "04"

	data = append(data, f(ID_CRC, checkSum(dataToCrc)))

	return serialize(data)
}

func (qr *PromptPayQR) GenerateBillPaymentPayload(billerID string, ref1 string, ref2 string, terminalID *string, amount *string) string {

	billerID = sanitizeTarget(billerID)

	var data []string
	data = append(data,
		f(ID_PAYLOAD_FORMAT, PAYLOAD_FORMAT_EMV_QRCPS_MERCHANT_PRESENTED_MODE),
		f(ID_POI_METHOD, ifThenElse(amount != nil, POI_METHOD_DYNAMIC, POI_METHOD_STATIC).(string)),
		f(ID_MERCHANT_INFORMATION_BOT_BILL_PAYMENT, serialize(
			[]string{
				f(BOT_ID_TAG30_AID, GUID_PROMPTPAY_BILL_PAYMENT),
				f(BOT_ID_TAG30_BILLER_ID, billerID),
				f(BOT_ID_TAG30_REF1, ref1),
				f(BOT_ID_TAG30_REF2, ref2),
			})),
		f(ID_TRANSACTION_CURRENCY, TRANSACTION_CURRENCY_THB),
	)

	if amount != nil {
		data = append(data, f(ID_TRANSACTION_AMOUNT, formatAmount(*amount)))
	}

	data = append(data, f(ID_COUNTRY_CODE, COUNTRY_CODE_TH))

	if terminalID != nil {
		data = append(data,
			f(ID_DATA_OBJECTS, serialize(
				[]string{
					f(BOT_ID_TAG62_TERMINAL_ID, *terminalID),
				})),
		)
	}

	dataToCrc := serialize(data) + ID_CRC + "04"

	data = append(data, f(ID_CRC, checkSum(dataToCrc)))

	return serialize(data)
}

func sanitizeTarget(value string) string {
	re := regexp.MustCompile(`[^0-9]`)
	value = re.ReplaceAllString(value, "")
	return value
}

func serialize(values []string) string {
	return strings.Join(values, "")
}

func f(id, value string) string {

	ext := "00" + strconv.Itoa(len(value))
	values := []string{id, ext[len(ext)-2:], value}
	value = strings.Join(values, "")
	return value
}

func formatTarget(value string) string {
	value = sanitizeTarget(value)
	if len(value) >= 13 {
		return value
	}

	re := regexp.MustCompile(`^0`)
	value = re.ReplaceAllString(value, "66")
	value = "0000000000000" + value

	return value[len(value)-13:]
}

func formatAmount(amount string) string {
	if f, err := strconv.ParseFloat(amount, 32); err == nil {
		value := fmt.Sprintf("%.2f", f)
		return value
	} else {
		panic(err)
	}

	return ""
}

func checkSum(value string) string {

	data := []byte(value)

	table := crc16.MakeTable(crc16.CRC16_CCITT_FALSE)
	crc := crc16.Checksum(data, table)
	h := fmt.Sprintf("%x", int(crc))

	return strings.ToUpper(h)

}

func ifThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

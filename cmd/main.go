/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"bytes"
	"fmt"
	"github.com/kazekim/promptpay-qr-go"
	"image"
	"image/png"
	"os"
)

func main() {

	qr, err := promptpayqr.QRForTargetWithAmount("0899999999","500" )

	if err != nil {
		panic(err)
	}

	byteToImage(*qr)

}

func byteToImage(imgByte []byte) {
	img, _, _ := image.Decode(bytes.NewReader(imgByte))

	//save the imgByte to file
	out, err := os.Create("./QRImg.png")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = png.Encode(out, img)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
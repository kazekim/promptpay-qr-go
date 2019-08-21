# promptpay-qr-go

GoLang Library to generate QR Code payload for PromptPay inspired from [dtinth/promptpay-qr](https://github.com/dtinth/promptpay-qr)

## Requirement
GoLang 1.12.x

## Install

    go get -u github.com/kazekim/promptpay-qr-go


## Implement

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

      // Image is return in []byte. You should convert to image by yourself.
	    qr, err := promptpayqr.QRForTargetWithAmount("0899999999","500" )

	    if err != nil {
		    panic(err)
	    }

      // I give an example of image convert here.
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

## Example of Output
You can support me by paying with this QR Code too. (LOL)
![example QR Code](https://github.com/kazekim/promptpay-qr-go/blob/master/cmd/QRImg.png?raw=true)
  

## Contributing
Everyone can contribute it. Feel free to improve it and make it better.

## License
The MIT License (MIT)

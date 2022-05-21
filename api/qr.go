package handler

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"log"
	"net/http"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

const ImageTemplate string = `<!DOCTYPE html>
<html lang="en"><head></head>
<body style="text-align:center">
<div>
	<img src="data:image/jpg;base64,{{.Image}}">
</div>
<strong>{{.Caption}}</strong>
</body>`

// http://localhost:3000/api/qr?lote=636
func QRHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("incoming request: qr")

	// Create the barcode
	// https://www.google.com/maps/dir/?api=1&destination=-34.3503079,-58.7630513&travelmode=driving
	qrCode, _ := qr.Encode("https://www.google.com/maps/dir/?api=1&destination=-34.3503079,-58.7630513&travelmode=driving", qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// encode the barcode as png
	//png.Encode(w, qrCode)

	writeImageWithTemplate(w, qrCode)

}

// Writeimagewithtemplate encodes an image 'img' in jpeg format and writes it into ResponseWriter using a template.
func writeImageWithTemplate(w http.ResponseWriter, img image.Image) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Fatalln("unable to encode image.")
	}

	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	if tmpl, err := template.New("image").Parse(ImageTemplate); err != nil {
		log.Println("unable to parse image template.")
	} else {
		data := map[string]interface{}{
			"Image":   str,
			"Caption": "Direcciones para ir al Lote 636",
		}

		if err = tmpl.Execute(w, data); err != nil {
			log.Println("unable to execute template.")
		}
	}
}

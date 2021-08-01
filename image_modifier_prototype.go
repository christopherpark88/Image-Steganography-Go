package main

import (
	//"fmt"
	"encoding/base64"
	"image"
	"os"
	//"io/ioutil"
	"log"
	_ "image/jpeg"
	//"net/http"
	//"strings"
	"image/color"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func check(err error) {  
	if err != nil {  
		panic(err)  
	}  
}

func main() {
	/* msg := "Hello, 世界"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println(string(decoded)) */
		// Read the entire file into a byte slice




		
	/* bytes, err := ioutil.ReadFile("./TestImage.jpg")
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	//fmt.Println(base64Encoding)	 */

	


	//Opens Image and creates blank rectangle to create output image
	reader, err := os.Open("./TestImage.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)

	if err != nil {
		log.Fatal(err)
	}

	size := m.Bounds().Size()  
    rect := image.Rect(0, 0, size.X, size.Y)  
    wImg := image.NewRGBA(rect)  

    for x := 0; x < size.X; x++ { 
        // and now loop thorough all of this x's y 
        for y := 0; y < size.Y; y++ { 
            pixel := m.At(x, y) 
            originalColor := color.RGBAModel.Convert(pixel).
                (color.RGBA) 
            // Offset colors a little, adjust it to your taste 
            r := float64(originalColor.R) * 0.92126 
            g := float64(originalColor.G) * 0.97152 
            b := float64(originalColor.B) * 0.90722 
            // average 
            grey := uint8((r + g + b) / 3) 
            c := color.RGBA{ 
                R: grey, G: grey, B: grey, A: originalColor.A, 
            } 
            wImg.Set(x, y, c) 
        } 
    } 
}
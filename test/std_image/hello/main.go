package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"os"
)

// https://www.devdungeon.com/content/working-images-go

// Generating an Image
func Generating() {

	// Create a blank image 10 pixels wide by 4 pixels tall : 10 * 4
	myImage := image.NewRGBA(image.Rect(0, 0, 10, 4))

	// You can access the pixels through myImage.Pix[i]
	// One pixel takes up four bytes/uint8. One for each: RGBA
	// So the first pixel is controlled by the first 4 elements
	// Values for color are 0 black - 255 full color
	// Alpha value is 0 transparent - 255 opaque
	myImage.Pix[0] = 255 // 1st pixel red
	myImage.Pix[1] = 0   // 1st pixel green
	myImage.Pix[2] = 0   // 1st pixel blue
	myImage.Pix[3] = 255 // 1st pixel alpha

	// myImage.Pix contains all the pixels
	// in a one-dimensional slice
	fmt.Println(myImage.Pix)
	fmt.Println(len(myImage.Pix))
	fmt.Println(myImage.Rect)

	// Stride is how many bytes take up 1 row of the image
	// Since 4 bytes are used for each pixel, the stride is
	// equal to 4 times the width of the image
	// Since all the pixels are stored in a 1D slice,
	// we need this to calculate where pixels are on different rows.
	fmt.Println(myImage.Stride) // 40 for an image 10 pixels wide  // 这就知道有多少个 像素点了

}

// Writing Image to File
func Writ2File() error {
	// Create a blank image 100x200 pixels
	myImage := image.NewRGBA(image.Rect(0, 0, 100, 200))

	// outputFile is a File type which satisfies Writer interface
	outputFile, err := os.Create("test.png")
	if err != nil {
		// Handle error
		return err

	}
	// Don't forget to close files
	defer outputFile.Close()

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	err = png.Encode(outputFile, myImage)
	if err != nil {
		return err
	}
	return nil
}

// Reading Image From File
func ReadFromFile() error {

	// Read image from file that already exists
	existingImageFile, err := os.Open("test.png")
	if err != nil {
		return err
	}
	defer existingImageFile.Close()

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	imageData, imageType, err := image.Decode(existingImageFile)
	if err != nil {
		return err
	}
	fmt.Println(imageData)
	fmt.Println(imageType)

	// We only need this because we already read from the file
	// We have to reset the file pointer back to beginning
	existingImageFile.Seek(0, 0) // 定位到文件开头

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		return err
	}
	fmt.Println(loadedImage)

	return nil

}

// Base64 Encoding Image
func Base64Encode() {
	// Create a blank image 10x20 pixels
	myImage := image.NewRGBA(image.Rect(0, 0, 10, 20))

	// In-memory buffer to store PNG image
	// before we base 64 encode it
	var buff bytes.Buffer

	// The Buffer satisfies the Writer interface so we can use it with Encode
	// In previous example we encoded to a file, this time to a temp buffer
	png.Encode(&buff, myImage)

	// Encode the bytes in the buffer to a base64 string
	encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())

	// You can embed it in an html doc with this string
	htmlImage := "<img src=\"data:image/png;base64," + encodedString + "\" />"
	fmt.Println(htmlImage)
}

// Base64 Encoding Image
func Base64EncodeFromFile() error {
	myImage, err := os.Open("/Users/apple/workspace_stariverpool/go-image/testdata/opencv/opencv-logo.png")
	if err != nil {
		return err
	}
	defer myImage.Close()

	imageData, imageType, err := image.Decode(myImage)
	if err != nil {
		return err
	}
	//fmt.Println(imageData)
	fmt.Println(imageType)

	// In-memory buffer to store PNG image
	// before we base 64 encode it
	var buff bytes.Buffer

	// The Buffer satisfies the Writer interface so we can use it with Encode
	// In previous example we encoded to a file, this time to a temp buffer
	png.Encode(&buff, imageData)

	// Encode the bytes in the buffer to a base64 string
	encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())

	// You can embed it in an html doc with this string
	htmlImage := "<img src=\"data:image/png;base64," + encodedString + "\" />"
	fmt.Println(htmlImage)

	return nil
}

func main() {
	//Generating()

	//Writ2File()

	//ReadFromFile()

	//Base64Encode()

	//Base64EncodeFromFile()

}

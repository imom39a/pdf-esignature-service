package pdfutils

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"pdf-esignature-server/blobstore"
	"strings"

	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

func SignFile(documentToSign string, base64Str string) string {

	tempFile, _ := blobstore.GetFileFromS3(documentToSign)

	pageNum := 3

	// read image file
	c := creator.New()

	img1, err := c.NewImageFromFile(TmpSignatureFile(base64Str))
	if err != nil {
		fmt.Println("Error opening image")
	}

	img2, err := c.NewImageFromFile(TmpSignatureFile(base64Str))
	if err != nil {
		fmt.Println("Error opening image")
	}

	img1.SetWidth(160)
	img1.SetHeight(50)
	img1.SetPos(50, 70)

	img2.SetWidth(160)
	img2.SetHeight(50)
	img2.SetPos(50, 170)

	// Read the input pdf file.
	f, err := os.Open(tempFile.Name())
	if err != nil {
		fmt.Println("Error opening PDF")
	}
	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		fmt.Println("Error creting PDF reader")
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		fmt.Println("Error getting page numbers")
	}

	fmt.Println("---> placing sign in pdf")

	// Load the pages.
	for i := 0; i < numPages; i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			fmt.Println("Error getting page numbers ... 1 !!!")
		}

		// Add the page.
		err = c.AddPage(page)
		if err != nil {
			fmt.Println("Error getting page numbers ... 2 !!!")
		}

		// If the specified page, or -1, apply the image to the page.
		if i+1 == pageNum || pageNum == -1 {
			_ = c.Draw(img1)
			_ = c.Draw(img2)
		}
	}

	signedFile, err := ioutil.TempFile("", "Signed-"+documentToSign)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("singed file ---> " + signedFile.Name())

	err = c.WriteToFile(signedFile.Name())

	return signedFile.Name()
}

func TmpSignatureFile(base64Data string) string {

	fmt.Println("---> TmpSignatureFile 1")
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Data))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println("base64toJpg", bounds, formatString)

	fmt.Println("---> TmpSignatureFile 2")
	signatureTmpFile, err := ioutil.TempFile("", "sig-*.jpeg")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("---> TmpSignatureFile 3")
	//Encode from image format to writer
	// f, err := os.OpenFile(signatureTmpFile.Name(), os.O_WRONLY|os.O_CREATE, 0777)
	// if err != nil {
	// 	log.Fatal(err)
	// 	panic(err)
	// }

	err = jpeg.Encode(signatureTmpFile, m, &jpeg.Options{Quality: 75})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return signatureTmpFile.Name()
}

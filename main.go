package main

import (
	"fmt"
	"goPDF/htmlParser"
	"goPDF/pdfGenerator"
)

type Data struct {
	Name string
}

func main() {

	h := htmlParser.New("tmp")
	wk := pdfGenerator.NewWkhtmlToPDF("tmp")

	dataHTML := Data{
		Name: "liz",
	}

	htmlGenerated, err := h.Create("templates/exemplo.html", dataHTML)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("html gerado", htmlGenerated)

	filePDFName, err := wk.Create(htmlGenerated)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("pdf gerado", filePDFName)
}

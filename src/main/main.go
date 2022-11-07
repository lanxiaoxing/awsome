package main

import (
	"fmt"
	"rsc.io/pdf"
)

func main() {
	fmt.Println("start analysis pdf....")
	file, err := pdf.Open("./src/main/2022.pdf")
	if err != nil {
		panic(err)
	}
	content := file.Page(2).Content()
	text := content.Text
	for _,v := range text {
		fmt.Print(v.S)
	}
}


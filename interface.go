package main

import "fmt"

type Document interface{
	Print() string
}

type TextDocument struct{}

type ImageDocument struct{}

func (t TextDocument) Print() string{
	return "hehe"
}

func (i ImageDocument) Print() string{
	return "haha"
}

func main(){
	text := TextDocument{}
	image := ImageDocument{}

	fmt.Println(text.Print())
	fmt.Println(image.Print())

}
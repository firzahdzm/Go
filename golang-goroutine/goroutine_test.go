package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func HelloWord(){
	fmt.Println("Hello World")
}

func TestGoroutine(t *testing.T){
	go HelloWord()
	fmt.Println("Hi")

	time.Sleep(time.Second)
}

func DisplayNumber(number int){
	fmt.Println(number)
}

func TestManyGoroutine(t *testing.T){
	for i := 0; i<10000; i++{
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}


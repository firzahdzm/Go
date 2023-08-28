package golanggoroutine

import (
	"fmt"
	"time"
	"testing"
	"strconv"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func(){
		time.Sleep(2 * time.Second)
		channel <- "Firza"
		fmt.Println("Done")
	}()

	res := <- channel
	fmt.Println(res)

	time.Sleep(4 * time.Second)
}

func GiveResponse(channel chan string)  {
	time.Sleep(2 * time.Second)
	channel <- "Firza"
}

func TestChannelParameter(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)

	go GiveResponse(channel)

	res := <- channel
	fmt.Println(res)

	time.Sleep(5 * time.Second)
}

func TestChannelBuffer(t *testing.T)  {
	channel := make(chan string, 5)
	defer close(channel)

	channel <- "Firza"
	channel <- "Firza"
	channel <- "Firza"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(len(channel))
	fmt.Println(cap(channel))

	fmt.Println("Done")
}

func TestChannelRange(t *testing.T){
	channel := make(chan string)

	go func(){
		for i := 0; i < 10; i++{
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		defer close(channel)
	}()

	for data := range channel{
		fmt.Println("menerima data ",data)
	}
}
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start!")

	for {
		fmt.Println("Running in the background")
		time.Sleep(time.Second * 2)
	}
}
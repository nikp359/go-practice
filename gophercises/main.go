package main

import (
	"fmt"
	"time"
)

func main() {
	goTimer()
}

func goTimer() {
	timer := time.NewTimer(5 * time.Second)
	//timeChan := make(chan string)
	for {
		select {
		case timeEnd := <-timer.C:
			fmt.Printf("time end: %v", timeEnd)
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("ping")
		}
	}
}

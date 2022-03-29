package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			select {
			case <-c:
				fmt.Println("Выхожу из программы")
				return
			default:
				fmt.Print(i * i)
				time.Sleep(time.Second)
				fmt.Print("\r     \r")
			}
		}
	}()

	wg.Wait()

}

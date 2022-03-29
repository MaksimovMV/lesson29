package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	wg := sync.WaitGroup{}
	wg.Add(2)
	intChan := scanner(&wg)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-c:
				fmt.Println("Выхожу из программы")
				return
			case num, ok := <-intChan:
				if !ok {
					return
				}
				fmt.Println("Квадрат числа:", num*num)
			}
		}
	}()

	wg.Wait()

}

func scanner(wg *sync.WaitGroup) chan int {
	s := bufio.NewScanner(os.Stdin)
	iChan := make(chan int)
	fmt.Println("Введите число или \"стоп\" для закрытия программы")
	go func() {
		for s.Scan() && s.Text() != "стоп" {
			num, err := strconv.Atoi(s.Text())
			if err != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			fmt.Println("Ввод:", num)
			iChan <- num
		}
		time.Sleep(time.Second)
		close(iChan)
		wg.Done()
	}()
	return iChan

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)
	intChan := scanner(&wg)

	go func() {
		defer wg.Done()
		for {
			num, ok := <-intChan
			if !ok {
				fmt.Println("До свидания")
				return
			}
			fmt.Println("Квадрат числа:", num*num)
		}
	}()

	wg.Wait()

}

func scanner(wg *sync.WaitGroup) chan int {
	s := bufio.NewScanner(os.Stdin)
	iChan := make(chan int)
	fmt.Println("Введите число или \"стоп\" для закрытия программы")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for s.Scan() {
			select {
			case <-c:
				fmt.Println("Выхожу из программы")
				break
			default:
				num, err := strconv.Atoi(s.Text())
				if err != nil {
					fmt.Println("Некорректный ввод")
					continue
				}
				fmt.Println("Ввод:", num)
				iChan <- num
			}
		}
		close(iChan)
		wg.Done()
	}()
	return iChan

}

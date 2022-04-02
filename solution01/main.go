package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	iChan := scanner(&wg)
	sqChan := square(&wg, iChan)
	multiplier(sqChan)
	wg.Wait()
}

func scanner(wg *sync.WaitGroup) chan int {
	s := bufio.NewScanner(os.Stdin)
	iChan := make(chan int)
	fmt.Println("Введите число или \"стоп\" для остановки программы")
	go func() {
		defer wg.Done()
		for s.Scan() && s.Text() != "стоп" {
			num, err := strconv.Atoi(s.Text())
			if err != nil {
				fmt.Println("Некорректный ввод:", err)
				continue
			}
			fmt.Println("Ввод:", num)
			iChan <- num
		}
		close(iChan)
	}()

	return iChan
}

func square(wg *sync.WaitGroup, iChan chan int) chan int {
	sqChan := make(chan int)

	go func() {
		defer wg.Done()
		for n := range iChan {
			sqChan <- n * n
			fmt.Println("Квадрат:", n*n)
		}
		close(sqChan)
	}()

	return sqChan
}

func multiplier(sqChan chan int) {
	for sq := range sqChan {
		fmt.Println("Произведение:", sq*2)
	}
}

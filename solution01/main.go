package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	iChan := scanner()
	sqChan := square(iChan)
	multiplier(sqChan)
}

func scanner() chan int {
	s := bufio.NewScanner(os.Stdin)
	iChan := make(chan int)
	fmt.Println("Введите число или \"стоп\" для остановки программы")
	go func() {
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

func square(iChan chan int) chan int {
	sqChan := make(chan int)

	go func() {
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

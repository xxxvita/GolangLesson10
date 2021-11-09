package main

import (
	"bufio"
	"fmt"
	"gb/GolangLesson10/fibo"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		err        error
		n          int64
		nFibonachi uint64
	)

	fmt.Printf("Программа вычисляет число Фибоначи по его порядковому номеру\n")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите номер числа Фибоначи:")
	for scanner.Scan() {

		str := strings.Trim(scanner.Text(), " ")
		n, err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			fmt.Printf("Ошибка ввода. Повторите попытку...\n")
			continue
		}

		//fmt.Printf("%p ,%d", cash, n)
		// Получение функции вычисления числа Фибоначи
		fFibonachi, err := fibo.GetFiboFunc(uint(n))
		if err != nil {
			fmt.Println(err)
			fmt.Printf("Введите номер числа Фибоначи:")

			continue
		}
		// Вычисление числа в последовательности Фибоначи с номером "n"
		// с учётом закешированных ранее результатов вычисления
		nFibonachi, err = fFibonachi()
		if err != nil {
			fmt.Println(err)
			fmt.Printf("Введите номер числа Фибоначи:")

			continue
		}

		// 0 1 2 3 5 8 13 21 34
		fmt.Printf("Число Фибоначи с номером %d равно %d\n", n, nFibonachi)
		fmt.Printf("Число итераций: %d!\n", fibo.CountStack)
		//fmt.Printf("%v", cash)

		fmt.Printf("Введите номер числа Фибоначи:")
	}
}

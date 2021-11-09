package fibo_test

import (
	"gb/GolangLesson10/fibo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibo(t *testing.T) {
	//assert := assert.New(t)

	fCalcFobonacci, err := fibo.GetFiboFunc(10)
	if err != nil {
		t.Fatal("Ошибка получения функции расчёта")
	}

	nFibonacci, err := fCalcFobonacci()
	if err != nil {
		t.Fatal("Ошибка при вычислении числа Фибоначчи")
	}

	assert.Equal(t, nFibonacci, 100, "Должно равняться 100")
}

package fibo

import (
	"errors"
	"fmt"
)

const Debug = 1

// Два близлежащих числа Фибоначи для вычисления на основе их следующего числа
type structFiboElement struct {
	fibo_2 uint64
	fibo_1 uint64
}

// Кеш-map для быстрого поиска ранее вычисленных значений чисел Фибоначи
// key - номер очередного числа Фибоначи
// value - структура, содержащая вычисленное значение самого числа Фибоначи и предыдущего
//         числа последовательности
type mapFibo map[uint]*structFiboElement

var cash mapFibo

// Число итераций, за которое определяется очередное число Фибоначи
// Для общей информации, чтобы видеть как помог кеш в виде числа новых вычислений
var CountStack uint

// Добавление ноовго значения в кеш
func (mf mapFibo) Append(item *structFiboElement) {
	mf[uint(len(mf))] = item
}

// Получение последнего сохранённого в кеше значения и его номер в последовательности
func (mf mapFibo) GetLasItem() (*structFiboElement, uint) {
	return mf[uint(len(mf)-1)], uint(len(mf) - 1)
}

// Правильная печать структуры structFiboElement (стандартная меняет последовательность филдов)
func (fe *structFiboElement) String() string {
	return fmt.Sprintf("{fibo_1:%d, fibo_2:%d}", fe.fibo_1, fe.fibo_2)
}

// Инициализация структур для хранения вычесленных последовательностей чисел Фибоначи
// Первые два элемента последовательности: 0, 1
func CashInit() {
	item0 := structFiboElement{
		fibo_1: 0,
		fibo_2: 0,
	}

	item1 := structFiboElement{
		fibo_1: 0,
		fibo_2: 1,
	}

	// Добавление очередного элемента последовательности в кеш
	cash = make(mapFibo)
	cash[0] = &item0
	cash[1] = &item1
}

// Получение следующего элемента в последовательности Фибоначи
func (c *structFiboElement) GetNext() (*structFiboElement, error) {
	if c == nil {
		err := errors.New("Не передан элемент цепочки Фибоначи для получения следующего")
		return nil, err
	}

	ch := new(structFiboElement)
	ch.fibo_1 = c.fibo_2
	ch.fibo_2 = c.fibo_1 + c.fibo_2

	return ch, nil
}

// получение текущего элемента из структуры хранения
func (c *structFiboElement) GetNumber() (uint64, error) {
	if c == nil {
		err := errors.New("Не передан элемент цепочки Фибоначи для получения содержащегося в нём числа")
		return 0, err
	}

	return c.fibo_1 + c.fibo_2, nil
}

// Я знаю о классическом примере вычисления чисел Фибоначи в виде суммы двух
// рекурсивных функций от F(n) и F(n-1).
// Мне он не нравится :-) Раз уж тренируем рекурсии, то я изобретаю свою (ближе к циклу).

// Вычисление n-ого числа из последовтаельности чисел Фибоначи
// С механизмом сохранения промежуточно вычисленных груп чисел.
// ch - описывает два последних вычисленных чисел Фибоначи
// n  - количество чисел, которые требуется вычислить дальше от ch
func Fibo(ch *structFiboElement, n uint) (uint64, error) {
	var (
		err  error
		item *structFiboElement
	)

	if ch == nil {
		return 0, errors.New("При вычислении нового числа в последовательности передано nil для текущего значения")
	}

	if n == 0 {
		return ch.GetNumber()
	} else {
		CountStack++
		// Добавление очередного элемента последовательности в кеш
		item, err = ch.GetNext()
		// мы не ожидаем провала при получении следующего значения в последовательности
		// даже переполнениек не вызывает остановку вычислений :-)
		if err != nil {
			return 0, err
		}

		cash.Append(item)
		// Получение следующего элемента последовательности
		return Fibo(item, n-1)
	}
}

// Возвращает фунцию, которая вычисляет число Фибоначи с заданным
// номером в последовательности number_element в зависимости
// от имеющихся на текущий момент вычисленных чисел
func GetFiboFunc(numberElement uint) (func() (uint64, error), error) {
	if numberElement < 0 {
		return nil, errors.New("Задано число не из натурального ряда")
	}

	var (
		numLastElem      uint // максимальный порядковый номер числа в сохранённой последовательности
		item             *structFiboElement
		countNextNumbers uint // Количество чисел последовательности, которые нужно довычислить
	)

	// Узнаю максимальный номер элемента последовательночти в кеше
	// и его значение
	item, numLastElem = cash.GetLasItem()

	if numberElement <= numLastElem {
		// Запрашиваемый элемент в цепочке уже был вычислен ранее
		numLastElem = numberElement
		item = cash[numLastElem]
		countNextNumbers = 0
	} else {
		countNextNumbers = numberElement - numLastElem
	}

	//fmt.Printf("Вычисление начнётся с числа под номером %d, будет произведено итераций %d\n", start_item, number_element-start_item)
	//fmt.Printf("Текущее значение стартового элемента %v\n", chain)
	return func() (uint64, error) {
		CountStack = 0
		// начальная точка в цепочке чисел Фибоначи и сколько ещё вычислить элементов
		return Fibo(item, countNextNumbers)
	}, nil
}

func init() {
	// Инициализация кеша
	CashInit()
}

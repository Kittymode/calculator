package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var arabToRome = map[int]string{ // карта для перевода арабских чисел в римские
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X",
}

var romeToArab = map[string]int{ // карта для перевода римских чисел в арабские
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func isRoman(input string) bool { // функция для проверки, является ли строка римским числом
	_, exist := romeToArab[input]
	return exist
}

func calc(a, b int, operator string) int { // функция для выполнения операции над двумя числами
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("ошибка: делить на 0 нельзя")
		}
		return a / b
	default:
		panic("ошибка: неизвестная операция")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение в формате: 1 + 1 или I + II")

	// Чтение строки целиком
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("ошибка при чтении ввода: %v", err))
	}

	// Убираем символ новой строки и лишние пробелы
	input = strings.TrimSpace(input)

	// Разделяем строку по пробелам для получения операндов и оператора
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("неверный формат ввода")
	}

	// Проверка, являются ли числа римскими
	isRomanInput := isRoman(parts[0]) && isRoman(parts[2])

	var a, b int
	var err1, err2 error
	operator := parts[1]

	if isRomanInput {
		// преобразуем римские числа в арабские
		a = romeToArab[parts[0]]
		b = romeToArab[parts[2]]
	} else {
		// преобразуем арабские числа
		a, err1 = strconv.Atoi(parts[0])
		b, err2 = strconv.Atoi(parts[2])
		if err1 != nil || err2 != nil {
			panic("ошибка: неверный формат чисел")
		}
	}

	// Проверка на диапазон чисел
	if a > 10 || b > 10 || a < 1 || b < 1 {
		panic("ошибка: числа должны быть в диапазоне от 1 до 10")
	}

	// Выполняем операцию
	result := calc(a, b, operator)

	// Выводим результат
	if isRomanInput {
		// Проверка на возможность представления результата в римских числах
		if result <= 0 {
			panic("ошибка: результат в римских числах должен быть положительным")
		}
		fmt.Println("Результат:", arabToRome[result])
	} else {
		fmt.Println("Результат:", result)
	}
}

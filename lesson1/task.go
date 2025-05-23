package lesson1

import (
	"fmt"
)

// Задача 1:
// ==============================================================
// Создайте переменную типа int и присвойте ей значение 10.
// Затем создайте еще одну переменную и присвойте ей значение 20.
// Выполните сложение этих двух переменных и выведите результат.
// ==============================================================

func Add(a int, b int) int {
	return a + b
}

// Задача 2:
// ==============================================================
// Создайте две переменные типа float64, присвойте им значения 3.14 и 2.71.
// Умножьте эти значения и выведите результат.
// ==============================================================

func Multiply(a, b float64) float64 {
	return a * b
}

// Задача 3:
// ==============================================================
// Создайте переменную типа string и присвойте ей строку "Hello, Go!".
// Затем выведите эту строку на экран.
// ==============================================================

func Greet() string {
	return "Hello, Go!"
}

// Задача 4:
// ==============================================================
// Создайте переменную типа rune и присвойте ей символ 'A'.
// Выведите этот символ и его числовое значение (код символа).
// ==============================================================

// Задача 5:
// ==============================================================
// Создайте переменную типа int со значением 10,
// затем создайте переменную типа float64 и присвойте ей значение переменной int,
// предварительно приведя типы.
// ==============================================================

func ConvertIntToFloat(a int) float64 {
	b := float64(a)
	return b
}

// Задача 6:
// ==============================================================
// Создайте переменную типа int со значением 5.
// Инкрементируйте её на 1 и затем декрементируйте на 1.
// Выведите результат после каждого изменения.
// ==============================================================

func IncrementAndDecrement(a int) (int, int) {
	a++
	a--
	return a, a
}

// Задача 7:
// ==============================================================
// Создайте переменную типа int, присвойте ей значение 7.
// Используя условный оператор if, верните bool значение,
// является ли число чётным или нечётным.
// ==============================================================

func CheckEvenOdd(a int) bool {
	if a%2 == 0 {
		return false
	} else {
		return true
	}
}

// Задача 8:
// ==============================================================
// Создайте программу, которая будет принимать номер дня недели (от 1 до 7) и
// выводить его название с использованием оператора switch.
// Если номер дня недели не входит в допустимый диапазон (1-7),
// программа должна вывести сообщение о неверном вводе.
// "Invalid", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"
// ==============================================================

func GetDayOfWeek(day int) string {
	switch day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid"
	}
}

// Задача 9:
// ==============================================================
// Создайте программу, которая использует type switch для проверки
// типа различных переменных и выводит соответствующие сообщения о типах этих переменных.
// ==============================================================

func CheckType(i interface{}) string {
	switch v := i.(type) {
	case int:
		return fmt.Sprintf("int: %d", v)
	case float64:
		return fmt.Sprintf("float64: %.6f", v) // Выравниваем формат до 6 знаков после точки
	case string:
		return fmt.Sprintf("string: %s", v)
	default:
		return "unknown type"
	}
}

// Задача 10:
// ==============================================================
// Создайте структуру Book, которая описывает книгу с полями:
// название, автор, цена и количество страниц.
// Затем создайте метод для этой структуры, который вычисляет скидку на книгу
// в зависимости от ее цены и выводит итоговую цену после скидки.
// ==============================================================

type Book struct {
	Title  string
	Author string
	Price  float64
}

func (b Book) DiscountedPrice(discount float64) float64 {
	discountt := b.Price * discount / 100
	finalPrice := b.Price - discountt
	fmt.Printf("Итоговая цена книги \"%s\" после скидки %.2f%%: %.2f\n", b.Title, discount, finalPrice)
	return finalPrice
}

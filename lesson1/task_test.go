package lesson1_test

import (
	tasks "tasks/lesson1"
	"testing"
)

func TestAdd(t *testing.T) {
	result := tasks.Add(10, 20)
	expected := 30

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	result := tasks.Multiply(3.14, 2.71)
	expected := 8.5094

	if result != expected {
		t.Errorf("expected %.4f, got %.4f", expected, result)
	}
}

func TestGreet(t *testing.T) {
	result := tasks.Greet()
	expected := "Hello, Go!"

	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestConvertIntToFloat(t *testing.T) {
	result := tasks.ConvertIntToFloat(10)
	expected := 10.0

	if result != expected {
		t.Errorf("expected %.1f, got %.1f", expected, result)
	}
}

func TestIncrementAndDecrement(t *testing.T) {
	resultX, resultY := tasks.IncrementAndDecrement(5)
	expected := 5

	if resultX != expected || resultY != expected {
		t.Errorf("expected %d, got %d and %d", expected, resultX, resultY)
	}
}

func TestCheckEvenOdd(t *testing.T) {
	result := tasks.CheckEvenOdd(7)
	expected := true

	if result != expected {
		t.Errorf("expected %t, got %t", expected, result)
	}
}

func TestGetDayOfWeek(t *testing.T) {
	result := tasks.GetDayOfWeek(3)
	expected := "Wednesday"

	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestCheckType(t *testing.T) {
	result := tasks.CheckType(42)
	expected := "int: 42"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}

	result = tasks.CheckType(3.14)
	expected = "float64: 3.140000"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}

	result = tasks.CheckType("Hello")
	expected = "string: Hello"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestDiscountedPrice(t *testing.T) {
	book := tasks.Book{Title: "Go Programming", Author: "John Doe", Price: 50.0}
	result := book.DiscountedPrice(10)
	expected := 45.0

	if result != expected {
		t.Errorf("expected %.2f, got %.2f", expected, result)
	}
}

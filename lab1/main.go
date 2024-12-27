package main

import (
	"errors"
	"fmt"
)

func hello(name string) string {
	return "Привет, " + name + "!"
}

func printEven(a, b int64) error {
	if a > b {
		return errors.New("левая граница диапазона больше правой")
	}

	for i := a; i <= b; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	return nil
}

func apply(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "−":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("действие не поддерживается")
	}
}

func main() {
	// hello
	fmt.Println(hello("Мир"))
	fmt.Println(hello("Go"))

	// printEven
	fmt.Println("Четные числа в диапазоне 1-12:")
	err := printEven(1, 12)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	fmt.Println("Четные числа в диапазоне 20-1:")
	err = printEven(20, 1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println("Четные числа в диапазоне -4-4:")
	err = printEven(-4, 4)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println("Четные числа в диапазоне 11-11:")
	err = printEven(11, 11)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println("Четные числа в диапазоне 12-12:")
	err = printEven(12, 12)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	// apply
	result, err := apply(3, 5, "+")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("3 + 5 =", result)
	}

	result, err = apply(7, 10, "*")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("7 * 10 =", result)
	}

	result, err = apply(3, 5, "#")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("3 # 5 =", result)
	}

	result, err = apply(10, 2, "/")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	result, err = apply(10, 0, "/")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("10 / 0 =", result)
	}

	result, err = apply(7, 3, "-")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("7 - 3 =", result)
	}
}

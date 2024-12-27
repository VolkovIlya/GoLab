package main

import (
	"errors"
	"fmt"
	"math"
)

// 1.1: Форматирование IP-адреса
func formatIP(ip [4]byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// 1.2: Список четных чисел в диапазоне
func listEven(a, b int) ([]int, error) {
	if a > b {
		return nil, errors.New("левая граница диапазона больше правой")
	}
	var result []int
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			result = append(result, i)
		}
	}
	return result, nil
}

// 2.1: Подсчет символов в строке
func countChars(s string) map[rune]int {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}
	return charCount
}

// 3.1: Структура "точка"
type Point struct {
	X float64
	Y float64
}

// 3.2: Структура "отрезок"
type Segment struct {
	Start Point
	End   Point
}

// 3.3: Метод для вычисления длины отрезка
func (s Segment) Length() float64 {
	dx := s.End.X - s.Start.X
	dy := s.End.Y - s.Start.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// 3.4: Структура "треугольник"
type Triangle struct {
	A Point
	B Point
	C Point
}

// 3.5: Структура "круг"
type Circle struct {
	Center Point
	Radius float64
}

// 3.6: Метод для вычисления площади треугольника
func (t Triangle) Area() float64 {
	a := Segment{t.A, t.B}.Length()
	b := Segment{t.B, t.C}.Length()
	c := Segment{t.C, t.A}.Length()
	p := (a + b + c) / 2
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

// 3.6: Метод для вычисления площади круга
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 3.7: Интерфейс "фигура"
type Shape interface {
	Area() float64
}

// 3.8: Функция для вывода площади
func printArea(s Shape) {
	result := s.Area()
	fmt.Printf("Площадь фигуры: %.2f\n", result)
}

// 4.1: Функция высшего порядка Map
func Map(slice []float64, f func(float64) float64) []float64 {
	result := make([]float64, len(slice))
	for i, value := range slice {
		result[i] = f(value)
	}
	return result
}

func main() {
	// Тест 1
	fmt.Println("Задание 1")
	ip := [4]byte{192, 168, 1, 100}
	fmt.Println("IP-адрес:", formatIP(ip))

	evenNumbers, err := listEven(1, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Четные числа:", evenNumbers)
	}

	evenNumbers, err = listEven(10, 1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Четные числа:", evenNumbers)
	}

	// Тест 2
	fmt.Println("\nЗадание 2: Карты")
	text := "hello world hello"
	charCounts := countChars(text)
	fmt.Println("Подсчет символов:", charCounts)

	// Тест 3
	fmt.Println("\nЗадание 3: Структуры, методы и интерфейсы")
	triangle := Triangle{
		A: Point{0, 0},
		B: Point{0, 3},
		C: Point{4, 0},
	}
	circle := Circle{
		Center: Point{0, 0},
		Radius: 5,
	}

	printArea(triangle)
	printArea(circle)

	segment := Segment{
		Start: Point{0, 0},
		End:   Point{3, 4},
	}
	fmt.Println("Длина отрезка:", segment.Length())

	// Тест 4
	fmt.Println("\nЗадание 4: Функциональное программирование")
	numbers := []float64{1, 2, 3, 4, 5}
	square := func(x float64) float64 {
		return x * x
	}

	squaredNumbers := Map(numbers, square)

	fmt.Println("Исходный срез:", numbers)
	fmt.Println("Срез после возведения в квадрат:", squaredNumbers)
}

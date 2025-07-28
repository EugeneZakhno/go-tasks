package main

import (
	"fmt"
	"sort"
)

func main() {

	//1.Верните несколько значений функции.
	println(twoValues(30, 2, "Саша", "Леша"))

	//2.Инициализируйте структуру
	var sa *Girl
	sa = new(Girl)
	sa.name = "Саша"
	sa.age = 9
	fmt.Printf("я %+v, мне %d лет\n", sa.name, sa.age)

	//2.Отсортируйте элементы слайса целых чисел.
	var slice = []int{5, 45, 50, 90, 45, -9}
	if !sort.IntsAreSorted(slice) { // утверждение = false     утверждение != true...
		sort.Ints(slice)
		fmt.Println("Sorted intSlice", slice)
	}
	n := len(slice)
	println(n)

	// 2*. Написать сортировку пузырьком
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < slice[j]-i-1; j++ {
			if slice[j] > slice[j+1] {
				// Меняем местами элементы, если они находятся в неправильном порядке
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	intSlice := []int{5, 3, 8, 1, 2}
	// Сортировка пузырьком
	s := len(intSlice)
	for i := 0; i < s-1; i++ {
		for j := 0; j < s-i-1; j++ {
			if intSlice[j] > intSlice[j+1] {
				// Меняем местами элементы, если они находятся в неправильном порядке
				intSlice[j], intSlice[j+1] = intSlice[j+1], intSlice[j]
			}
		}
	}
	fmt.Println("Отсортированный массив:", intSlice)
}

// 1.Верните несколько значений функции.
func twoValues(a, b int, s, l string) (int, int, string) {
	var conc = s + l
	return a + b, a * b, conc
}

type Girl struct {
	name string
	age  int
}

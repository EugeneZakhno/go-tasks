package main

func main() {

	//1.Верните несколько значений функции.
	println(twoValues(30, 2))
}

// 1.Верните несколько значений функции.
func twoValues(a, b int) (int, int, int, int) {
	return a + b, a * b, a - b, a / b
}

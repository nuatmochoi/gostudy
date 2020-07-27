package main

import (
	"fmt"
	"math"
	"strings"
)

func Sqrt(x float64) float64 {
	z := 1.0 // == var z float64 = 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func mul(a, b int) int {
	return a * b
}

func lenAndUppser(name string) (length int, uppercase string) {
	defer fmt.Println("I'm Done") // func이 끝나고 실행
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	// for _, number := range numbers {
	// 	fmt.Print(number)
	// }
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

func repeat(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println(mul(2, 2))

	totalLength, upperName := lenAndUppser("choi")
	fmt.Println(totalLength, upperName)
	repeat("choi", "seong", "woo", "hi", "wifi")

	fmt.Println(superAdd(1, 2, 3, 4, 5, 6))
}

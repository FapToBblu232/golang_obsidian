package main

import "fmt"

func main() {
	costi := make([][2]int, 0)
	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			elem := [2]int{i, j}
			costi = append(costi, elem)
		}
	}
	double := make([][2][2]int, 0)
	for i := 0; i < 27; i++ {
		for j := i + 1; j < 28; j++ {
			pair := [2][2]int{costi[i], costi[j]}
			double = append(double, pair)
		}
	}
	count := 0
	for _, elem := range double {
		if elem[0][0] == elem[1][0] || elem[0][0] == elem[1][1] || elem[0][1] == elem[1][0] || elem[0][1] == elem[1][1] {
			count++
		}
	}
	fmt.Println((count - 84) / 21)
}

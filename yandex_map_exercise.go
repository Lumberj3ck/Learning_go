package main

import "fmt"

func get_order_price(order []string, pricelist map[string]int) (orderPrice int) {
	for _, product := range order {
		orderPrice += pricelist[product]
	}
	return orderPrice
}

func initialiseTask() {
	var pricelist map[string]int
	pricelist = make(map[string]int)
	pricelist["butter"] = 50
	pricelist["milch"] = 100
	pricelist["brot"] = 80

	price := get_order_price([]string{"milch", "brot"}, pricelist)
	fmt.Println(price)
}

type Ind int

func pairOfNumbers(A []int, k int) (Ind, Ind) {
	fract := make(map[int]Ind)
	for idx1, num := range A {
		rest := k - num
		if idx2, ok := fract[rest]; ok {
			fmt.Println("Values", A[idx1], A[idx2])
			return Ind(idx1), idx2
		}
		fract[num] = Ind(idx1)
	}
	return 0, 0
}

func RemoveDuplicates(input []string) []string {
	words := make(map[string]int)
	updated_input := make([]string, 0)
	for _, word := range input {
		if _, ok := words[word]; ok == false {
			updated_input = append(updated_input, word)
			words[word] = 1
		}
	}
	return updated_input
}

func main() {
	// A := []int{5, 2, 4, 1, 3}
	// k := 6
	// fmt.Println(A)
	// idx1, idx2 := pairOfNumbers(A, k)
	// fmt.Println(idx1, idx2)
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	i := RemoveDuplicates(input)
	fmt.Println(i)
}

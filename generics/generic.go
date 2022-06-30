package main

import "fmt"

type Number interface {
	int8 | int32 | int64 | float32 | float64
}

func newGenricFunc[age Number](myage age) {
	val := float32(myage) + 1
	fmt.Println((val))
}

func BubbleSort[N Number](input []N) []N {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}

		}
	}

	return input
}

func main() {
	var age int64 = 28
	var age2 float32 = 28.5

	newGenricFunc(age)
	newGenricFunc(age2)

	list := []int32{4, 3, 2, 1, 7}
	listFloat := []float32{4.5, 2.8, 1.5, 3.5}

	sortedList := BubbleSort(list)
	fmt.Println(sortedList)

	sortedFloatList := BubbleSort(listFloat)
	fmt.Println(sortedFloatList)
}

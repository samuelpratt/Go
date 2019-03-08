package main

import "fmt"

func main() {
	var b = []int{1, 2, 3, 4, 5}

	var sum = sum(b)

	fmt.Println("The sum is", sum)
}

func sum(arrayToSum []int) int {
	if len(arrayToSum) == 1 {
		return arrayToSum[0]
	}
	return arrayToSum[0] + sum(arrayToSum[1:])

}

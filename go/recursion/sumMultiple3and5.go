package main

import "fmt"

//Find the sum of all the multiples of 3 or 5 below 1000. 233168

func findSum(num int) int{
	sum3 := findSumMultipleOf3(num-1)
	sum4 := findSumMultipleOf5(num-1)
	return sum3+sum4
}

func findSumMultipleOf3(num int) int{
	var sum int
	if num == 0 {
		return sum
	} else {
		if num%3 == 0 {
			sum += num
		}
	}
	num--
	return sum + findSumMultipleOf3(num)
}

func findSumMultipleOf5(num int) int {
	var sum int
	if num == 0 {
		return sum
	} else {
		if num %5 == 0  && num % 3 != 0 {
			sum += num
		}
	}
	num--
	return sum + findSumMultipleOf5(num)
}

func main() {
	n := 10000
	fmt.Print(findSum(n))

}

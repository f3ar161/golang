package main

import "fmt"

func main() {

	result := sumProcess(6)
	result2 := channelsF(5, 2)
	fmt.Println("result: ", result)
	fmt.Println("result2: ", result2)

}

func sumProcess(value int) int {
	var items = make(map[int]int)
	var finalItems = make(map[int]int)
	items[1] = 1
	items[2] = 2
	items[3] = 3
	items[4] = 4
	limit := 4

	for i := 1; i <= value; i++ {
		sumTemp := 0
		for j := i; j <= limit; j++ {
			sumTemp += items[j]
		}
		limit++
		finalItems[i] = sumTemp
		items[limit] = sumTemp
		if limit == value {
			break
		}
	}

	return finalItems[len(finalItems)]
}

func channelsF(n int, k int) int {
	var channelsN []int

	// generate channels
	for i := 1; i <= n; i++ {
		channelsN = append(channelsN, i)
	}

	for {

		fmt.Println("result: ", channelsN)

		channelsN = append(channelsN[:k-1], channelsN[k:]...)
		if len(channelsN) == 2 {
			break
		}

	}

	return channelsN[1]

}

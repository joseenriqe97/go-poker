package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	TestAlgorithm([]int{14, 5, 4, 2, 3})
}

func TestAlgorithm(cardsItem []int) (bool, string) {
	var (
		response string
		count    int
	)
	//TODO: It remains to validate the length of the array and that it is not less than 2
	sort.Slice(cardsItem, func(i, j int) bool {
		return cardsItem[i] < cardsItem[j]
	})

	count = 1

	for x := range cardsItem {
		i := sort.Search(len(cardsItem), func(i int) bool { return cardsItem[i] >= cardsItem[x]+1 })
		if i < len(cardsItem) && cardsItem[i] == cardsItem[x]+1 {
			count = count + 1
		} else if count < 4 {
			count = 1
		}
	}
	if count == 4 && cardsItem[len(cardsItem)-1] == 14 && cardsItem[0] == 2 {
		response = arrayToString(cardsItem, ",")
		return true, response
	}
	if count == 5 {
		response = arrayToString(cardsItem, ",")
		return true, response
	}
	return false, arrayToString(cardsItem, ",")
}

func arrayToString(item []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(item), " ", delim, -1), "[]")
}

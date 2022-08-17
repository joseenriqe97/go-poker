package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	TestAlgorithm([]int{14, 5, 4, 2, 3, 1})
}

func TestAlgorithm(cardsItem []int) (bool, string) {
	var (
		response string
		count    int
	)
	sort.Slice(cardsItem, func(i, j int) bool {
		return cardsItem[i] < cardsItem[j]
	})

	if len(cardsItem) > 7 {
		return false, arrayToString(cardsItem, ",")
	}

	if cardsItem[0] < 2 {
		return false, arrayToString(cardsItem, ",")
	}

	if cardsItem[len(cardsItem)-1] > 14 {
		return false, arrayToString(cardsItem, ",")
	}

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

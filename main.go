package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Print(TestAlgorithm([]int{2, 3, 4, 5, 6}))
}

func TestAlgorithm(cardsItem []int) (bool, string) {
	var (
		auxP     int
		response string
	)
	//Para romper referencia
	cards := make([]int, len(cardsItem))
	copy(cards, cardsItem)

	deckOfCards := make(map[string][]int)
	brokeSequence := false

	sort.Slice(cards, func(i, j int) bool {
		return cards[i] < cards[j]
	})
	for x := range cards {

		if !brokeSequence && x < len(cards)-1 && cards[x]-cards[x+1] == -1 {

			if len(deckOfCards["firstHand"]) > 0 {
				deckOfCards["firstHand"] = append(deckOfCards["firstHand"], cards[x+auxP])

			} else {
				deckOfCards["firstHand"] = append(deckOfCards["firstHand"], cards[x], cards[x+1])
				auxP = x + 1
			}

		} else {
			brokeSequence = true
		}
		if brokeSequence && x < len(cards)-1 && cards[x]-cards[x+1] == -1 {
			auxP = x + 1
			if len(deckOfCards["secondHand"]) > 0 && x < len(cards)-1 {
				deckOfCards["secondHand"] = append(deckOfCards["secondHand"], cards[auxP+1])
			} else {
				deckOfCards["secondHand"] = append(deckOfCards["secondHand"], cards[x], cards[x+1])

			}
			if x == len(cards)-1 {
				deckOfCards["secondHand"] = append(deckOfCards["secondHand"], cards[x])
			}
		}
	}
	if len(deckOfCards["firstHand"]) == 5 {
		response = arrayToString(cardsItem, ",", false)
		return true, response
	}

	if len(deckOfCards["secondHand"]) == 5 {
		response = arrayToString(cardsItem, ",", false)
		return true, response
	}
	if len(deckOfCards["firstHand"]) == 4 && cards[len(cards)-1] == 14 {

		/* if deckOfCards["firstHand"][0] == 2 {
			deckOfCards["firstHand"] = append(deckOfCards["firstHand"][:], cards[len(cards)-1])
		} */

		response = arrayToString(cardsItem, ",", false)
		return true, response
	}
	response = arrayToString(cardsItem, ",", false)

	return false, response
}

func arrayToString(item []int, delim string, asc bool) string {
	if asc {
		sort.Slice(item, func(i, j int) bool {
			return item[i] > item[j]
		})
	}
	return strings.Trim(strings.Replace(fmt.Sprint(item), " ", delim, -1), "[]")
}

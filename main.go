package main

import (
	"fmt"
	"sort"
)

func main() {
	/* cartas := [5]string{10, 20, 30, 40, 50}
	fmt.Println("hola") */

	testAlgorithm([]int{7, 3, 2})
}

func testAlgorithm(cards []int) {
	var auxP int
	deckOfCards := make(map[string][]int)
	brokeSequence := false
	//Ordenamos el array

	sort.Slice(cards, func(i, j int) bool {
		return cards[i] < cards[j]
	})
	for x := range cards {
		// fmt.Print(cards[x+1] - cards[x])
		/* if x == 0 && cards[x+1]-cards[x] == 1 {
			// aux = []int{cards[x], cards[x+1]}
			aux = append(aux, cards[x], cards[x+1])
		} */
		if !brokeSequence && x < len(cards)-1 && cards[x]-cards[x+1] == -1 {
			/* if deckOfCards["1MANO"] == nil {
				deckOfCards["1MANO"] = make([]int, cards[x], cards[x+1])
			} */

			if len(deckOfCards["1MANO"]) > 0 {
				// fmt.Println("entramops", cards[x], cards[x+auxP])
				deckOfCards["1MANO"] = append(deckOfCards["1MANO"], cards[x+auxP])

			} else {
				deckOfCards["1MANO"] = append(deckOfCards["1MANO"], cards[x], cards[x+1])
				auxP = x + 1
			}
			// fmt.Print("entramos     ")

			fmt.Println(deckOfCards["1MANO"])
		} else {
			brokeSequence = true
		}
		if brokeSequence && x < len(cards)-1 && cards[x]-cards[x+1] == -1 {
			if len(deckOfCards["2MANO"]) > 0 && x < len(cards)-1 {
				// fmt.Println("entramops", cards[x], cards[x+auxP])
				deckOfCards["2MANO"] = append(deckOfCards["2MANO"], cards[auxP+1])
				auxP = x + 1
			} else {
				deckOfCards["2MANO"] = append(deckOfCards["2MANO"], cards[x], cards[x+1])
				auxP = x + 1
			}
			if x == len(cards)-1 {
				deckOfCards["2MANO"] = append(deckOfCards["2MANO"], cards[x])
			}
		}
		// fmt.Println(auxP)
		/* if x > 0 && x != len(cards)-1 {
			if aux[len(aux)-1] != cards[x] {
				aux = append(aux, cards[x])
			}
		}
		if x == len(cards)-1 && cards[x-1]-cards[x] == -1 {
			aux = append(aux, cards[x])
		} */

		// fmt.Print(cards[x])
	}

	/* if aux[len(aux)-1]-aux[0] != aux[len(aux)-1]-1 {
		fmt.Print("false")
	} */
	fmt.Print(deckOfCards)

}

// func find()

/* 2 5 7 8 9 10 11 */

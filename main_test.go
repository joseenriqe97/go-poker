package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlgorithmScenery(t *testing.T) {

	res, cards := TestAlgorithm([]int{2, 3, 4, 5, 6})
	assert.Equal(t, res, true)
	assert.Equal(t, cards, "2,3,4,5,6")

	res, cards = TestAlgorithm([]int{14, 5, 4, 2, 3})
	assert.Equal(t, res, true)
	assert.Equal(t, cards, "2,3,4,5,14")

	res, cards = TestAlgorithm([]int{7, 7, 12, 11, 3, 4, 14})
	assert.Equal(t, res, false)
	assert.Equal(t, cards, "3,4,7,7,11,12,14")

	res, cards = TestAlgorithm([]int{7, 3, 2})
	assert.Equal(t, res, false)
	assert.Equal(t, cards, "2,3,7")

}

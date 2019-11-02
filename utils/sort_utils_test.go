package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	elements := []int{9,8,7,6,5,4}

	sortedElements := BubbleSort(elements)

	assert.NotNil(t, sortedElements)

	assert.EqualValues(t, 6, len(sortedElements))
	assert.EqualValues(t, 4, sortedElements[0])
	assert.EqualValues(t, 5, sortedElements[1])
	assert.EqualValues(t, 6, sortedElements[2])
	assert.EqualValues(t, 7, sortedElements[3])
	assert.EqualValues(t, 8, sortedElements[4])
	assert.EqualValues(t, 9, sortedElements[5])
}
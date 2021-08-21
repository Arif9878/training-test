package main_test

import (
	"testing"

	q4 "github.com/Arif9878/stockbit-test/question-4"
	"gotest.tools/v3/assert"
)

func Test_SortingString(t *testing.T) {
	output := q4.SortString("stockbit")

	assert.Equal(t, "bcikostt", output)
}

package main_test

import (
	"testing"

	q3 "github.com/Arif9878/stockbit-test/question-3"
	"gotest.tools/v3/assert"
)

func Test_FindFirstStringInBracket(t *testing.T) {

	output := q3.FindFirstStringInBracket("(stockbit)")

	assert.Equal(t, "stockbit", output)
}

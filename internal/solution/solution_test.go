package solution

import (
	"testing"

	"math/rand"

	"github.com/stretchr/testify/assert"
)

func TestBasicCase(t *testing.T) {
	input := []int{1, 3, 6, 4, 1, 2}
	expected := 5

	actual := Solution(input)

	assert.Equal(t, expected, actual)
}

func TestOrdered(t *testing.T) {
	input := []int{1, 2, 3}
	expected := 4

	actual := Solution(input)

	assert.Equal(t, expected, actual)
}

func TestNegatives(t *testing.T) {
	input := []int{-1, -3}
	expected := 1

	actual := Solution(input)

	assert.Equal(t, expected, actual)
}

func TestEquals(t *testing.T) {
	input := []int{1, 1}
	expected := 2

	actual := Solution(input)

	assert.Equal(t, expected, actual)
}

func TestOneElement(t *testing.T) {
	input := []int{1}
	expected := 2

	actual := Solution(input)

	assert.Equal(t, expected, actual)
}

func getRandomN() int {
	return rand.Intn(1_000_000) - 1_000_000
}

const MAX = 100_000

func TestBigInput(t *testing.T) {

	//N is an integer within the range [1..100,000];
	//each element of array A is an integer within the range [−1,000,000..1,000,000].
	var input []int
	for i := 0; i < MAX; i++ {
		input = append(input, 1)
	}
	expected := 2

	actual := Solution(input)

	assert.Equal(t, expected, actual)
}

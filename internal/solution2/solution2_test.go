package solution2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicCase1(t *testing.T) {
	X := 100 * SECOND

	fmt.Println("-----------------------------")
	fmt.Printf("X: %v\n", X)
	fmt.Println("-----------------------------")

	expected := "1m40s"

	actual := Solution2(X)

	assert.Equal(t, expected, actual)
}

func TestBasicCase2(t *testing.T) {
	X := 7263

	fmt.Println("-----------------------------")
	fmt.Printf("X: %v\n", X)
	fmt.Println("-----------------------------")

	expected := "2h2m"

	actual := Solution2(X)

	assert.Equal(t, expected, actual)
}

func TestBasicCase3(t *testing.T) {
	X := 5 * SECOND

	fmt.Println("-----------------------------")
	fmt.Printf("X: %v\n", X)
	fmt.Println("-----------------------------")

	expected := "5s"

	actual := Solution2(X)

	assert.Equal(t, expected, actual)
}

func TestZero(t *testing.T) {
	X := 0 * SECOND

	fmt.Println("-----------------------------")
	fmt.Printf("X: %v\n", X)
	fmt.Println("-----------------------------")

	expected := "0s"

	actual := Solution2(X)

	assert.Equal(t, expected, actual)
}

func TestOneWeek(t *testing.T) {
	X := WEEK

	fmt.Println("-----------------------------")
	fmt.Printf("X: %v\n", X)
	fmt.Println("-----------------------------")

	expected := "1w"

	actual := Solution2(X)

	assert.Equal(t, expected, actual)
}

func TestRounding(t *testing.T) {
	X := WEEK + DAY + HOUR + MINUTE

	fmt.Println("-----------------------------")
	fmt.Printf("X: %v\n", X)
	fmt.Println("-----------------------------")

	expected := "1w2d"

	actual := Solution2(X)

	assert.Equal(t, expected, actual)
}

func TestOneWeekOneDay(t *testing.T) {
	X := WEEK + DAY

	fmt.Println("-----------------------------")
	fmt.Printf("X: %v\n", X)
	fmt.Println("-----------------------------")

	expected := "1w1d"

	actual := Solution2(X)

	assert.Equal(t, expected, actual)
}

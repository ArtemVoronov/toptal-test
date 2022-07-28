package solution3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicCase1(t *testing.T) {
	S := "300.01"
	B := []string{"300.00", "200.00", "100.00"}

	fmt.Println("-----------------------------")
	fmt.Printf("S: %v, B: %v\n", S, B)
	fmt.Println("-----------------------------")

	expected := []string{"150.00", "100.00", "50.01"}

	actual := Solution3(S, B)

	assert.Equal(t, expected, actual)
}

func TestBasicCase2(t *testing.T) {
	S := "1.00"
	B := []string{"0.05", "1.00"}

	fmt.Println("-----------------------------")
	fmt.Printf("S: %v, B: %v\n", S, B)
	fmt.Println("-----------------------------")

	expected := []string{"0.05", "0.95"}

	actual := Solution3(S, B)

	assert.Equal(t, expected, actual)
}

package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Add(2, 2) = %d; want %d", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

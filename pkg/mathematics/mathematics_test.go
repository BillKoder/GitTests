package mathematics

import (
	"testing"
)

func TestSum(t *testing.T) {
	if Sum([]int{1, 1}) != 2 {
		t.Error("Expected 1 + 1 to equal 2")
	}
}

func TestMultiply(t *testing.T) {
	if Multiply([]int{3, 3}) != 9 {
		t.Errorf("Expected 2 * 2 to equal 4")
	}
}

func TestDifference(t *testing.T) {
	if Difference(10, 5) != 5 {
		t.Errorf("Expected 10 - 5 to equal 5")
	}
}

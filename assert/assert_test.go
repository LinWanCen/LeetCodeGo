package assert

import (
	"testing"
)

func TestAssertEqualArr_value(t *testing.T) {
	AssertEqualArr(t, []int{
		11, 222, 3, 4,
	}, []int{
		11, 22, 33, 5,
	})
}

func TestAssertEqualArr_len(t *testing.T) {
	AssertEqualArr(t, []int{
		11, 222, 3,
	}, []int{
		1, 2,
	})
}

func TestAssertEqualMatrix_value(t *testing.T) {
	AssertEqualMatrix(t, [][]int{
		{11, 22},
		{11, 222, 3},
	}, [][]int{
		{11, 222},
		{11, 22, 4},
	})
}

func TestAssertEqualMatrix_len(t *testing.T) {
	AssertEqualMatrix(t, [][]int{
		{11, 222},
		{11, 222, 3},
	}, [][]int{
		{11, 222},
		{1, 2},
	})
}

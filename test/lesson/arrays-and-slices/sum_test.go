package arrays_and_slices

import (
	"fmt"
	testingTools "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	arraysAndSlices "github.com/starsky-mik/kgl-go-learing/lesson/arrays-and-slices"
	"testing"
)

type SumCase struct {
	array  []int
	result int
}

func TestSum(t *testing.T) {
	testCases := []SumCase{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{0, 0, 0}, 0},
		{[]int{-1, 0, 1}, 0},
		{[]int{}, 0},
	}

	for _, testCase := range testCases {
		t.Run("check arrays_and_slices.Sum function", func(t *testing.T) {
			testingTools.AssertEquals(t, arraysAndSlices.Sum(testCase.array), testCase.result)
		})
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arraysAndSlices.Sum([]int{1, 2, 3, 4, 5})
	}
}

func ExampleSum() {
	sum := arraysAndSlices.Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(sum)
	// Output 15
}
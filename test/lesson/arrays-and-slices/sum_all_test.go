package arrays_and_slices

import (
	"fmt"
	testingTools "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	arraysAndSlices "github.com/starsky-mik/kgl-go-learing/lesson/arrays-and-slices"
	"testing"
)

type SumAllCase struct {
	arrays [][]int
	result []int
}

func TestSumAll(t *testing.T) {
	testCases := []SumAllCase{
		{[][]int{{1, 2}, {2, 3}}, []int{3, 5}},
		{[][]int{{0, 0}, {0, 0}}, []int{0, 0}},
		{[][]int{{1, -1}, {-1, 1}}, []int{0, 0}},
		{[][]int{{1, 2}, {}}, []int{3, 0}},
		{[][]int{{}, {}}, []int{0, 0}},
	}

	for _, testCase := range testCases {
		t.Run("check arrays_and_slices.SumAll function", func(t *testing.T) {
			testingTools.AssertDeepEquals(t, arraysAndSlices.SumAll(testCase.arrays...), testCase.result)
		})
	}

	t.Run("check arrays_and_slices.SumAll function with empty slices list", func(t *testing.T) {
		actual := arraysAndSlices.SumAll()
		var expected []int

		testingTools.AssertDeepEquals(t, actual, expected)
	})
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arraysAndSlices.SumAll([]int{1, 2}, []int{2, 3}, []int{3, 4})
	}
}

func ExampleSumAll() {
	sum := arraysAndSlices.SumAll([]int{1, 2}, []int{2, 3}, []int{3, 4})
	fmt.Println(sum)
	// Output [3 5 7]
}
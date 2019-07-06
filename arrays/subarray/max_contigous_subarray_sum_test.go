package subarray

import (
	"github.com/RanjeetKaur17/go-coding-practice/arrays"
	"testing"
)

func Test_MaxSubArraySum(t *testing.T) {
	S := MaxSubArraySum([]int{0})

	if S != 0 {
		t.Fail()
	}

	S = MaxSubArraySum([]int{-1})

	if S != -1 {
		t.Fail()
	}

	S = MaxSubArraySum([]int{-2,1,-3,4,-1,2,1,-5,4})

	if S != 6 {
		t.Fail()
	}
}

func Test_MaxSumSubArray(t *testing.T) {
	S := MaxSumSubArray([]int{0})

	if !arrays.ArrayEquals(S, []int{0}) {
		t.Fail()
	}

	S = MaxSumSubArray([]int{-1})

	if !arrays.ArrayEquals(S, []int{-1}) {
		t.Fail()
	}

	S = MaxSumSubArray([]int{-2,1,-3,4,-1,2,1,-5,4})

	if !arrays.ArrayEquals(S, []int{4,-1,2,1}) {
		t.Fail()
	}
}
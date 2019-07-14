package tree

import "testing"

func Test_IsValidBST(t *testing.T) {
	isValid := IsValidBST(SampleBST())

	if !isValid {
		t.Fail()
	}

	isValid = IsValidBST(SampleBinaryTree())

	if isValid {
		t.Fail()
	}

	isValid = IsValidBST(nil)

	if !isValid {
		t.Fail()
	}

	isValid = IsValidBST(&BinaryTree{value: 1})

	if !isValid {
		t.Fail()
	}
}
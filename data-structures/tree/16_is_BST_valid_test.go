package tree

import "testing"

func Test_IsValidBST(t *testing.T) {
	isValid := IsValidBST(SampleBST())

	if !isValid {
		t.Fail()
	}

	isValid = IsValidBST(SampleTree())

	if isValid {
		t.Fail()
	}

	isValid = IsValidBST(nil)

	if !isValid {
		t.Fail()
	}

	isValid = IsValidBST(&Tree{value:1})

	if !isValid {
		t.Fail()
	}
}

func SampleBST() *Tree {
	return &Tree{
		value:10,
		left:&Tree{
			value:5,
			left:&Tree{
				value:2,
			},
			right:&Tree{
				value:7,
				left:&Tree{
					value:6,
				},
				right:&Tree{
					value:8,
				},
			},
		},
		right:&Tree{
			value:15,
			left:&Tree{
				value:12,
			},
			right:&Tree{
				value:20,
			},
		},
	}
}
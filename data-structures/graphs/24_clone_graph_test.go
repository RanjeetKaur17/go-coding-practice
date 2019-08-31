package graphs

import "testing"

func Test_CloneGraph(t *testing.T) {
	var root *UGraphNode

	clone := CloneGraph(root)
	if clone != nil {
		t.Fail()
	}

	root = &UGraphNode{label:1}

	clone = CloneGraph(root)
	if clone == nil || clone.label != 1 || len(clone.neighbours) != 0 {
		t.Fail()
	}

	root = &UGraphNode{label:1, neighbours:[]*UGraphNode{{label:2}, {label:3}}}

	clone = CloneGraph(root)
	if clone == nil || clone.label != 1 || len(clone.neighbours) != 2 ||
		clone.neighbours[0].label != 2 || len(clone.neighbours[0].neighbours) != 0 ||
		clone.neighbours[1].label != 3 || len(clone.neighbours[1].neighbours) != 0 {
		t.Fail()
	}
}
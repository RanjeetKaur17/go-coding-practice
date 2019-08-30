package graphs

type node struct {
	val string
	neighbors []*node
	depth int
}

func WordLadderLength(start, end string, dict []string) int {
	m := map[string]*node{}

	m[start] = &node{start, nil, 0}
	m[end] = &node{end, nil, 0}

	for _, d := range dict {
		m[d] = &node{d, nil, 0}
	}

	for _, d1 := range m {
		for _, d2 := range m {
			if match(d1.val, d2.val) {
				m[d1.val].neighbors = append(m[d1.val].neighbors, m[d2.val])
			}
		}
	}

	setDepth(m[start], 1)

	return m[end].depth
}

func setDepth(n *node, depth int) {
	if n.depth == 0 {
		n.depth = depth
		for _, n2 := range n.neighbors {
			setDepth(n2, depth+1)
		}
	}
}

func match(A,B string) bool {
	mismatch := 0
	for i := range A {
		if A[i] != B[i] {
			mismatch++
		}
	}

	if mismatch == 1 {
		return true
	}
	return false
}

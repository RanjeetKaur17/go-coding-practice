package graphs

type node struct {
	val string
	neighbors []*node
	depth int
}

func WordLadderLength(start, end string, dict []string) int {

	m := buildGraph(append(dict, start, end))
	setDepth(m[start], 1)

	return m[end].depth
}

func WordLadder(start, end string, dict []string) [][]string {

	m := buildGraph(append(dict, start, end))
	setDepth(m[start], 1)

	if m[end].depth != 0 {
		return paths(m, start, end)
	}

	return [][]string{}
}

func paths(m map[string]*node, start , end string) [][]string {
	if start == end {
		return [][]string{{start}}
	}
	res := [][]string{}
	for _, n := range m[end].neighbors {
		if n.depth == m[end].depth-1 {
			sub_res := paths(m, start, n.val)
			for _, s := range sub_res {
				res = append(res, append(s, m[end].val))
			}
		}
	}

	return res
}

func buildGraph(list []string) map[string]*node {
	m := map[string]*node{}

	for _, d := range list {
		m[d] = &node{d, nil, 0}
	}

	for _, d1 := range m {
		for _, d2 := range m {
			if match(d1.val, d2.val) {
				m[d1.val].neighbors = append(m[d1.val].neighbors, m[d2.val])
			}
		}
	}
	return m
}

func setDepth(n *node, depth int) {
	if n.depth == 0 || depth < n.depth {
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

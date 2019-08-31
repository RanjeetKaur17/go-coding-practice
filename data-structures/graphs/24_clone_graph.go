package graphs

/*
Problem Statement:
Clone an undirected graph. Each node in the graph contains a label and a list of its neighbors.

Sample Input:

 */
func CloneGraph(node *UGraphNode) *UGraphNode {
	if node == nil {
		return nil
	}

	m := map[int]*UGraphNode{}
	Parse(node, m)

	newM := map[int]*UGraphNode{}

	for k := range m {
		newM[k] = &UGraphNode{label:k}
	}

	for k, v := range m {
		for _, n := range v.neighbours {
			newM[k].neighbours = append(newM[k].neighbours, newM[n.label])
		}
	}

	return newM[node.label]
}

func Parse(node *UGraphNode, m map[int]*UGraphNode) {
	if _, ok := m[node.label]; !ok {
		m[node.label] = node
		for _, n := range node.neighbours {
			Parse(n, m)
		}
	}
}
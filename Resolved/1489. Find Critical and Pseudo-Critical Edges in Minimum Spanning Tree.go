func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
    // Use Kruskal's algo to find MST
    // Wef can sort edges first, since that algo needs sorted edges to use in greedy approach

    // first store indicies before sorting
    sortedEdges := make([][]int, len(edges))

    for i, edge := range edges {
        sortedEdges[i] = []int{edge[0], edge[1], edge[2], i}
    }
    sort.Slice(sortedEdges, func(i, j int) bool {
        return sortedEdges[i][2] < sortedEdges[j][2]
    })
    //fmt.Println(sortedEdges)
    result := make([][]int, 2)

    mst, mstWeight := findMst(n, sortedEdges, -1, -1)
    //fmt.Println("Found first MST", mst, mstWeight)
    // now find critical edges
    // critical edges are in mst, and removing them will increase mst weight
    result[0] = make([]int, 0)
    for _, mstEdgeIndex := range mst {
        _, newWeight := findMst(n, sortedEdges, mstEdgeIndex, -1)
        //fmt.Println(mstEdgeIndex, newWeight, newMst)
        if newWeight > mstWeight {
            result[0] = append(result[0], sortedEdges[mstEdgeIndex][3])
        }
    }

    // pseudo-critical edge are: not in critical edge list, and give the same weight as original MST
    result[1] = make([]int, 0)
    for i := range sortedEdges {
        isCritical := false
        for _, critEdge := range result[0] {
            if sortedEdges[i][3] == critEdge {
                isCritical = true
                break
            }
        }
        if !isCritical {
            _, newWeight := findMst(n, sortedEdges, -1, i)
            if newWeight == mstWeight {
                result[1] = append(result[1], sortedEdges[i][3])
            }
        }
    }

    return result
}

// returns MST of given graph and its weight
func findMst(n int, edges [][]int, excludedEdgeIndex int, includedEdgeIndex int) ([]int, int) {
    // result will be list of edge indices, size of not more than n - 1
    result := make([]int, 0)

    // required variables for DSU
    representative, size := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		representative[i] = i
	}

    weight := 0


    // force to add this edge
    if includedEdgeIndex != -1 {
        result = append(result, includedEdgeIndex)
        edge := edges[includedEdgeIndex]
        combine(representative, size, edge[0], edge[1])
        weight += edge[2]
    }


    // i - num of edge in input list of edges
    i := 0
    for len(result) < n - 1 {
        if i == len(edges) {
            // we couldn't create connected graph, so no MST
            return []int{}, math.MaxInt32
        }
        if i == excludedEdgeIndex || i == includedEdgeIndex {
            // includedEdgeIndex is alredy included
            i++
            continue
        }
        edge := edges[i]
        x := find(representative, size, edge[0]);
        y := find(representative, size, edge[1]);
        //fmt.Println("Representative for", i, edge, x, y)

        if x != y {
            // no cycle!
            //fmt.Println("Added", i)
            result = append(result, i)
            combine(representative, size, edge[0], edge[1])
            weight += edge[2]
        } else {
            //fmt.Println("Cycle, not added", i)
        }

        i++
    }

    return result, weight
}

// DSU find
func find(representative []int, size []int, u int) int {
	if u == representative[u] {
		return u
	} else {
		representative[u] = find(representative, size, representative[u])
		return representative[u]
	}
}

// DSU union
func combine(representative []int, size []int, u, v int) {
	u = find(representative, size, u)
	v = find(representative, size, v)

	if u == v {
		return
	} else {
		if size[u] > size[v] {
			representative[v] = u
		} else if size[u] < size[v] {
			representative[u] = v
		} else {
			representative[v] = u
			size[u]++
        }
	}
}

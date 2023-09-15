// https://leetcode.com/problems/min-cost-to-connect-all-point
func minCostConnectPoints(points [][]int) int {

    n := len(points)

    sortedEdges := make([][]int, 0)
    for i, p1 := range points {
        for j := i + 1; j < len(points); j++ {
            p2 := points[j]
            dist := abs(p1[0]-p2[0]) + abs(p1[1] - p2[1])
            sortedEdges = append(sortedEdges, []int{i, j, dist})
            sortedEdges = append(sortedEdges, []int{j, i, dist})
        }
    }
    
    sort.Slice(sortedEdges, func(i, j int) bool {
        return sortedEdges[i][2] < sortedEdges[j][2]
    })

    //fmt.Println(sortedEdges)
    
    // use Kruskal’s Minimum Spanning Tree (MST) Algorithm
    _, mstWeight := findMst(n, sortedEdges)

    return mstWeight
}

// returns MST of given graph and its weight
func findMst(n int, edges [][]int) ([]int, int) {
    // result will be list of edge indices, size of not more than n - 1
    result := make([]int, 0)

    // required variables for DSU
    representative, size := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		representative[i] = i
	}

    weight := 0


    // i - num of edge in input list of edges
    i := 0
    for len(result) < n - 1 {
        if i == len(edges) {
            // we couldn't create connected graph, so no MST
            return []int{}, math.MaxInt32
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

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

// https://leetcode.com/problems/shortest-cycle-in-a-graph/
func findShortestCycle(n int, edges [][]int) int {
	links := make(map[int][]int)

	for _, link := range edges {
		links[link[0]] = append(links[link[0]], link[1])
		links[link[1]] = append(links[link[1]], link[0])
	}
	min := math.MaxInt32
	notInCycleNodes := make(map[int]bool)
    skipNodes := make(map[int]bool) // skip nodes that were in cycle
	for i := 0; i < n; i++ {
        if !notInCycleNodes[i] {
            res := bfs(i, n, links, &notInCycleNodes, &skipNodes)
            if res < min {
                min = res
                skipNodes[i] = true
            }
        }
	}
	if min == math.MaxInt32 {
		return -1
	}
	return min
}

func bfs(i, n int, links map[int][]int, notInCycleNodes *map[int]bool, skipNodes *map[int]bool) int {
	result := math.MaxInt32
	visited := make(map[int]bool)

	visitedEdges := map[int]bool{}

	queue := []int{i}

	levels := make(map[int]int)

	visited[i] = true
	levels[i] = 1

    processedNodes := []int{}
    cycleFound := false

	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		//fmt.Println("Popped", popped)
        processedNodes = append(processedNodes, popped)

		for _, child := range links[popped] {
            hash := popped * 1e6 + child
			if !(*skipNodes)[child] && !visitedEdges[hash] { // not a simple link between child and popped
                //fmt.Println(popped, child)
				if !visited[child] {
                    //fmt.Println(popped, "=>", child)
					visited[child] = true
					visitedEdges[hash] = true
                    reverseHash := child * 1e6 + popped
					visitedEdges[reverseHash] = true
					levels[child] = levels[popped] + 1
					queue = append(queue, child)
				} else  {
                    //fmt.Println(i, "Already visited", child, popped, levels[child], levels[popped])
                    currDist := levels[popped] + levels[child] - 1
                    if currDist < result {
                        result = currDist
                    }
                    cycleFound = true
                }
			}
		}

	}
    if !cycleFound {
        // here we don't have a cycle
        for _, k := range processedNodes {
            (*notInCycleNodes)[k] = true
        }
    }
	return result
}

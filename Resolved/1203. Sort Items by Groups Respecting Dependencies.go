// https://leetcode.com/problems/sort-items-by-groups-respecting-dependencies/
func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
    // create graph for groups and items(nodes)
    // graph of nodes will give us order of items in result, and graph of groups will ensure same group items will come together
    // let's assign separate group for each item with group = -1
    for i := 0; i < n; i++ {
        if group[i] == -1 {
            group[i] = m
            m++
        }
    }

    graphGroup := make([][]int, m)
    for i := range graphGroup {
        graphGroup[i] = make([]int, 0)
    }
    graphNode := make([][]int, n)
    for i := range graphNode {
        graphNode[i] = make([]int, 0)
    }
    // fill in graphs for groups and nodes
    for i := 0; i < n; i++ {
        if len(beforeItems[i]) > 0 {
            for _, parent := range beforeItems[i] {
                graphNode[parent] = append(graphNode[parent], i)

                if group[i] != group[parent] {
                    graphGroup[group[parent]] = append(graphGroup[group[parent]], group[i])
                }
            }
        } 
    }

    sortedGraphGroup, correct := topologicalSort(m, graphGroup)
    if !correct {
        //fmt.Println("Cycle in groups")
        return []int{}
    }
    sortedGraphNode, correct := topologicalSort(n, graphNode)
    if !correct {
        //fmt.Println("Cycle in items", graphNode)
        return []int{}
    }
    
    result := []int{}
    for i := m-1; i >= 0; i--  {
        grp := sortedGraphGroup[i] 
        for j := n -1; j >= 0; j-- {
            node := sortedGraphNode[j] 
            if group[node] == grp {
                result = append(result, node)
            }
        }
    }

    return result
}

// returns ordered list of nodes, and false, if cycle found
func topologicalSort(count int, graph [][]int) ([]int, bool) {
    visited := map[int]bool{}
    stack := []int{}

    // returns false if cycle found
    var sortUtil func(node int, cycleCheck *map[int]bool) bool
    sortUtil = func(node int, cycleCheck *map[int]bool) bool {        
        (*cycleCheck)[node] = true
        visited[node] = true

        for _, child := range graph[node] {
            if (*cycleCheck)[child] {
                //fmt.Println("Cycle end found", node, "=>", child)
                return false
            }
            if !visited[child] {
                if !sortUtil(child,cycleCheck) {
                    //fmt.Println("Cycle found", node, "=>", child)
                    return false
                }
            }
        }
        (*cycleCheck)[node] = false
        stack = append(stack, node)
        return true
    }

    for i := range graph {
        if !visited[i] {
            cycleCheck := map[int]bool{}
            if !sortUtil(i, &cycleCheck) {
                return []int{}, false
            }
        }
    }

    /*result := make([]int, count)

    for i, v := range stack {
        result[count - i - 1] = v
    }*/
    return stack, true
}

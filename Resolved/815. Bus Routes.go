// https://leetcode.com/problems/bus-routes/
func numBusesToDestination(routes [][]int, source int, target int) int {
    // Need to create graph between buses, not points
    // And then BFS to find min length
    if source == target {
        return 0
    }
    buses := make(map[int][]int)
    linkMaps := make([]map[int]bool, len(routes))

    sourceBuses := make([]int, 0)
    targetBuses := make(map[int]bool)
    for bus, route := range routes {
        linkMaps[bus] = make(map[int]bool)
        for i := 0; i < len(route); i++ {
            dest := route[i]
            buses[dest] = append(buses[dest], bus)

            if dest == source {
                sourceBuses = append(sourceBuses, bus)
            }
            if dest == target {
                targetBuses[bus] = true
            }
        }
    }

   
    for _, routeBuses := range buses {
        for i := 0; i < len(routeBuses); i++ {
            for j := 0; j < len(routeBuses); j++ {
                if i != j {
                    //fmt.Println("Conn", i, j)
                    linkMaps[routeBuses[i]][routeBuses[j]] = true
                    linkMaps[routeBuses[j]][routeBuses[i]] = true
                }
            }
        }
    }

    links := make([][]int, len(routes))
    for i := 0; i < len(routes); i++ {
        links[i] = make([]int, 0)
        for j := 0; j < len(routes); j++ {
            if linkMaps[i][j] {
                 links[i] = append(links[i], j)   
            }
        }   
    }

    /*fmt.Println("Source", sourceBuses)
    fmt.Println("Target", targetBuses)

    fmt.Println(links)*/
    result := -1
    for _, sourceBus := range sourceBuses {
        currRes, found := bsf(links, targetBuses, sourceBus)
        if found {
            if result == -1 || result > currRes {
                result = currRes
            }
        }
    }

    return result
}

func bsf(links [][]int, targetBuses map[int]bool, bus int) (int, bool) {
    visited := make(map[int]bool)
    level := make(map[int]int)

    visited[bus] = true
    level[bus] = 1

    queue := []int{bus}

    for len(queue) > 0 {
        popped := queue[0]
        queue = queue[1:]

        if targetBuses[popped] {
            return level[popped], true
        }

        for _, nextBus := range links[popped] {
            if !visited[nextBus] {
                visited[nextBus] = true
                level[nextBus] = level[popped] + 1
                
                queue = append(queue, nextBus)
            }
        }
    }

    return 0, false
}

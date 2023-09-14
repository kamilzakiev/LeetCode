// https://leetcode.com/problems/reconstruct-itinerary/
func findItinerary(tickets [][]string) []string {
    
    links := make(map[string][]string)
    used := make(map[string]map[int]bool)
    for _, ticket := range tickets {
        links[ticket[0]] = append(links[ticket[0]], ticket[1])
        used[ticket[0]] = make(map[int]bool)
    }

    circuit := []string{}

    stack := []string{"JFK"}
    current := "JFK"

    // Hierholzer’s Algorithm
    for len(stack) > 0 {
        min := ""
        minIndex := -1
        for i, v := range links[current] { // can be optimized
            if !used[current][i] {
                if min == "" || min > v {
                    min = v
                    minIndex = i
                }
            }
        }        
        
        // if there is still edge from current
        if min != "" {
            stack = append(stack, current)
            used[current][minIndex] = true
            current = min
        } else {
            circuit = append(circuit, current)
            current = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
    }

    result := []string{}

    for i := len(circuit)-1; i >= 0; i-- {
        result = append(result, circuit[i])
    }


    return result
}

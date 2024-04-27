// https://leetcode.com/problems/freedom-trail/
func findRotateSteps(ring string, key string) int {
    places := make(map[byte][]int)

    for i := 0; i < len(ring); i++ {
        places[ring[i]] = append(places[ring[i]], i)
    }

    //fmt.Println(places)

    memo := make([][]int, len(key)) //min number of steps for i-th character of key starting at j ring position
    for i := range memo {
        memo[i] = make([]int, len(ring))
    }


    var dfs func(i, j int) int

    // we start at position j of ring and find min number of steps to spell i-th character of key
    dfs = func(i, j int) int {
        if i == len(key) {
            return 0
        }
        if memo[i][j] > 0 {
            return memo[i][j]
        }
        
        result := math.MaxInt32
        // all positions in ring of next character from key
        for _, charJ := range places[key[i]] {
            fmt.Println(key[i], charJ)
            // 1. First we need to rotate to current character in ring
            steps := diff(j, charJ, len(ring))
            // 2. Then spell
            steps++
            // 3. Do same for the next character starting at charJ position of ring            
            steps += dfs(i+1, charJ)

            result = min(result, steps)
        }
        memo[i][j] = result
        //fmt.Println(i, j, result)
        return result        
    }

    return dfs(0, 0)
}

// returns number of steps to rotate from i to j, l is len(ring)
func diff(i, j, l int) int {
    return min(min(abs(i-j), abs(l+i-j)), abs(l+j-i))
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}


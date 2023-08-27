// https://leetcode.com/problems/frog-jump/
func canCross(stones []int) bool {
    mapStones := map[int]bool{}

    // cache of stoneNum + k
    memo := map[int]map[int]bool{}

    for _, stone := range stones {
        mapStones[stone] = true
        memo[stone] = map[int]bool{}
    }

    // check if we can go to end if make jump of k length from currStone
    var canJump func(k int, currStone int) bool

    canJump = func(k int, currStone int) bool {
        //fmt.Println("Checking for", currStone, k)

        if savedResult, exists := memo[currStone][k]; exists {
            return savedResult
        }
        nextStone := k + currStone

        if !mapStones[nextStone] {
            memo[currStone][k] = false
            return false
        }
        // reached last stone
        if nextStone == stones[len(stones)-1] {
            //fmt.Println("Reached last stone", currStone, k)
            memo[currStone][k] = true
            return true
        }
        result := false
        if k > 1 {
            result = canJump(k-1, nextStone)
        }
        if !result && k > 0 {
            result = canJump(k, nextStone)
        }
        if !result {
            result = canJump(k+1, nextStone)
        }
        //fmt.Println(result, currStone, k)
        memo[currStone][k] = result

        return result
    }
    return canJump(1, 0)

}

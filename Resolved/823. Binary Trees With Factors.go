// https://leetcode.com/problems/binary-trees-with-factors/
func numFactoredBinaryTrees(arr []int) int {
    const mod = 1e9+7
    mapI := make(map[int]bool)

    for _, v := range arr {
        mapI[v] = true
    }

    children := make(map[int]map[int]bool)

    for _, x := range arr {
        children[x] = make(map[int]bool)
        for _, y := range arr {
            if !children[x][y] && x % y == 0 && mapI[x/y] && y <= x/y { // add less divisor instead of both
                children[x][y] = true
            }
        }
    }
    
    memo := make(map[int]int)

    var numOfTrees func(root int) int

    numOfTrees = func(root int) int {
        if memo[root] > 0 {
            return memo[root]
        }
        result := 1
        if children[root] != nil {
            for child := range children[root] {
                multiplier := 2
                if root/child == child {
                    multiplier = 1
                }
                ch1, ch2 := numOfTrees(child), numOfTrees(root/child)
                result = int(int64(result) + (int64(multiplier) * int64(ch1) * int64(ch2)) % mod)
            }
        }
        memo[root] = result
        return result
    }

    result := 0
    for _, v := range arr {
        count := numOfTrees(v)
        result = int((int64(result) + int64(count)) % mod)
    }

    return result
}

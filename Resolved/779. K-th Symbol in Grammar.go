// https://leetcode.com/problems/k-th-symbol-in-grammar
func kthGrammar(n int, k int) int {
    if n == 1 {
        return 0
    }
    var iter func(int, int, int) int

    iter = func(nodeVal, lastLevelNodeCount, kIndex int) int {
        left := kIndex <= lastLevelNodeCount / 2
        //fmt.Println(nodeVal, lastLevelNodeCount, kIndex, left)
        nextNodeVal := 0
        if nodeVal == 0 && !left || nodeVal ==1 && left {
            nextNodeVal = 1
        }
        if lastLevelNodeCount == 2 {
            return nextNodeVal
        }
        kIndex = (kIndex-1) % (lastLevelNodeCount / 2) + 1
        lastLevelNodeCount = lastLevelNodeCount / 2
        return iter(nextNodeVal, lastLevelNodeCount, kIndex)
    }

    return iter(0, int(math.Pow(2, float64(n-1))), k)
}

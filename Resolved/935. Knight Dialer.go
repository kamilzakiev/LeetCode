// https://leetcode.com/problems/knight-dialer/description/
func knightDialer(n int) int {
    const mod = 1e9+7
    next := map[int][]int{1:[]int{6, 8}, 2:[]int{7, 9}, 3:[]int{4, 8}, 4:[]int{3, 9, 0}, 6:[]int{1, 7, 0}, 7:[]int{2, 6}, 8:[]int{1, 3}, 9:[]int{2, 4}, 0:[]int{4, 6}}
    
    memo := map[int]int{}
    var iter func(i int, cell int) int 
    iter = func(i int, cell int) int {
        if i == n {
            return 1
        } else {
            hash := cell * 10000 + i
            if memo[hash] > 0 {
                return memo[hash]
            }
            result := int64(0)
            for _, nextCell := range next[cell] {
                result += int64(iter(i+1, nextCell))
            }
            memo[hash] = int(result % mod)
            return memo[hash]
        }
    }
    result := 0
    for cell := 0; cell <= 9; cell ++ {
        result = int((int64(result) + int64(iter(1, cell))) % mod)
    }
    return result
}

// https://leetcode.com/problems/minimum-penalty-for-a-shop/
func bestClosingTime(customers string) int {
    n := len(customers)
    p1 := make([]int, n+1)
    p2 := make([]int, n+1)

    p1[0] = 0 // not needed, default is 0 already
    p1Comb := 0
    for i := 1; i <= n; i++ {
        if p[i-1] == 'N' {
            p1Comb++
        }
        p1[i] = p1Comb
    }

    p2[n] = 0
    p2Comb := 0

    for i := n-1; i >= 0; i-- {
        if p2[i] == 'Y' {
            p2Comb++
        }
        p2[i] = p2Comb
    }

    minVal := math.MaxInt32

    for i :=0; i <= n; i++ {
        sum := p1[i] + p2[i]
        if sum < minVal {
            minVal = sum
        }
    }

    for i :=0; i <= n; i++ {
        if 
    } 
}

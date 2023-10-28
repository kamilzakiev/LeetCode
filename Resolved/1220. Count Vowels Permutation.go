// https://leetcode.com/problems/count-vowels-permutation/
func countVowelPermutation(n int) int {
    const mod = 1e9+7
    next := make([][]int, 5) // define as next[i][j] = 0 or 1, meaning that after vowel i can or not come vowel j
    for i := range next {
        next[i] = make([]int, 5)
    }
    next[0][1] = 1 //Each vowel 'a' may only be followed by an 'e'.
    next[1][0], next[1][2] = 1, 1 //Each vowel 'e' may only be followed by an 'a' or an 'i'.
    next[2][0], next[2][1], next[2][3], next[2][4] = 1, 1, 1, 1 // Each vowel 'i' may not be followed by another 'i'.
    next[3][2], next[3][4] = 1, 1 // Each vowel 'o' may only be followed by an 'i' or a 'u'.
    next[4][0] = 1 // Each vowel 'u' may only be followed by an 'a'.

    dp := make([][]int, n+1) // dp[i][j] returns how many string exists of length i ending with j vowel
    
    for i := range dp {
        dp[i] =  make([]int, 5)
    }

    for j := 0; j < 5; j++ {
        dp[1][j] = 1
    }
    for i := 2; i <= n; i++ {
        for j := 0; j < 5; j++ {
            for prev := 0; prev < 5; prev++ {
                dp[i][j] = int((int64(dp[i][j]) + int64(dp[i-1][prev] * next[prev][j])) % mod)
            }
        }
    }
    result := 0
    for j := 0; j < 5; j++ {
        result = int((int64(result) + int64(dp[n][j])) % mod)
    }

    return result
    
}

// https://leetcode.com/problems/unique-binary-search-trees/
func numTrees(n int) int {
    // there will be n nodes
    // root value can be from 1 to n
    // if root = i, then i-1 nodes on the right subtree, n-i nodes on the left subtree
    dp := make([]int, n+1)
    dp[0] = 1
    dp[1] = 1
    for count := 2; count <= n; count++ {
        for i := 1; i <= count; i++ { // what value is in the root if we have 'count' number of  nodes
            dp[count] += dp[i-1] * dp[count-i]
        }
    }
    return dp[n]
}

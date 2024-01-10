// https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) int {
    links := make(map[int][]int)

    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        if node.Left != nil {
            links[node.Val] = append(links[node.Val], node.Left.Val)
            links[node.Left.Val] = append(links[node.Left.Val], node.Val)
            dfs(node.Left)
        }
        if node.Right != nil {
            links[node.Val] = append(links[node.Val], node.Right.Val)
            links[node.Right.Val] = append(links[node.Right.Val], node.Val)
            dfs(node.Right)
        }
    }
    dfs(root)
    //fmt.Println(links)
    queue := []int{start}
    visited := make(map[int]bool)
    visited[start] = true
    length := make(map[int]int)
    length[start] = 0
    result := 0
    for len(queue) > 0 {
        popped := queue[0]
        queue = queue[1:]
        for _, next := range links[popped] {
            if !visited[next] {
                visited[next] = true
                queue = append(queue, next)
                length[next] = length[popped]+1
                result = max(result, length[next])
            }
        }
    }
    return result
}

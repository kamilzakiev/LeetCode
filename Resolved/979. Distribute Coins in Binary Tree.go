// https://leetcode.com/problems/distribute-coins-in-binary-tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
    // idea is to count how many coins will go to left and right node from current
    // to understand that, we need to count how many coins and nodes are on left and right sides of the subtree of this node
    result := 0

    var sumCoin func(node *TreeNode) (int, int) // node count, coin sum

    sumCoin = func(node *TreeNode) (int, int) {
        if node == nil {
            return 0, 0
        }

        leftCount, leftCoinSum := sumCoin(node.Left)
        rightCount, rightCoinSum := sumCoin(node.Right)

        result += abs(leftCount - leftCoinSum) + abs(rightCount - rightCoinSum)

        return leftCount + rightCount + 1, leftCoinSum + rightCoinSum + node.Val
    }

    sumCoin(root)

    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

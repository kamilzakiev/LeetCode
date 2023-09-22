// https://leetcode.com/problems/flatten-binary-tree-to-linked-list
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    if root == nil {
        return
    }
    var parentToRight *TreeNode // end of left branch, it will be parent for right tree branch
    if root.Left != nil {
        flatten(root.Left)
        curr := root.Left 
        for curr != nil {
            parentToRight = curr
            curr = curr.Right
        }
    }
    if root.Right != nil {
        flatten(root.Right)
    }
    if root.Left != nil && root.Right != nil {
        prevRight := root.Right
        root.Right = root.Left
        root.Left = nil
        parentToRight.Right = prevRight
    } else if root.Left != nil {
        root.Right = root.Left
        root.Left = nil
    } // if left == nil, do nothing
}

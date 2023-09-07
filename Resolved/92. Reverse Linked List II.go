// https://leetcode.com/problems/reverse-linked-list-ii
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    result := head
    var leftEnd *ListNode // most right not changed node on left part
    curr := head
    for i := 0 ; i < left - 1; i++ {
        leftEnd = curr
        curr = curr.Next
    }
    midStart := curr // most left node in changed part
    prev := curr
    curr = curr.Next
    for i := left; i < right; i++ {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next        
    }
    midEnd := prev
    midStart.Next = curr // 2 => 5
    if leftEnd != nil {
        leftEnd.Next = midEnd // 1 => 4 
    } else {
        result = midEnd
    }

    return result
}

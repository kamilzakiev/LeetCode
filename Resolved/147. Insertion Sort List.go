// https://leetcode.com/problems/insertion-sort-list/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    sorted := false
    for !sorted {
        curr := head
        var relocate *ListNode
        for curr != nil {
            if curr.Next != nil && curr.Val > curr.Next.Val {
                relocate = curr.Next
                curr.Next = curr.Next.Next
                break
            } else {
                curr = curr.Next
            }
        }
        if relocate != nil {
            if head.Val > relocate.Val {
                relocate.Next = head
                head = relocate
            } else {
                curr = head
                prev := head
                for curr.Val <= relocate.Val {
                    prev = curr
                    curr = curr.Next
                }
                prev.Next = relocate
                relocate.Next = curr
            }
        } else {
            sorted = true
        }
    }
    return head
}

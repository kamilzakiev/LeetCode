// https://leetcode.com/problems/rotate-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    curr := head
    count := 1

    for curr.Next != nil {
        count++
        curr = curr.Next
    }
    curr.Next = head

    unlinkNum := count - k % count - 1 // index of element that should have Next = nil, index starting by 0

    i := 0
    curr = head
    for i < unlinkNum {
        curr = curr.Next
        i++
    }
    head = curr.Next
    curr.Next = nil
    return head
}

// https://leetcode.com/problems/split-linked-list-in-parts
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
    curr := head
    count := 0
    
    for curr != nil {
        count++
        curr = curr.Next
    }

    remCount := count

    result := []*ListNode{}

    curr = head
    var prev *ListNode
    for i := 0; i < k; i++ {        
        elemCount := remCount / (k - i)

        if remCount % (k - i) > 0 {
            elemCount++
        }

        if prev != nil {
            prev.Next = nil
        }

        if elemCount == 0 {
            result = append(result, nil)
        } else {           
            partHead := curr
            for j := 0; j < elemCount; j++ {
                prev = curr
                curr = curr.Next
                remCount--
            }
            result = append(result, partHead)
        }
    }
    return result
}

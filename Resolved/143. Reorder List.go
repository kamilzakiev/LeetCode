// https://leetcode.com/problems/reorder-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    list := make([]*ListNode, 0)
    curr := head
    for curr != nil {
        list = append(list, curr)
        curr = curr.Next
    }
    prev := head
    for len(list) > 0 {
        first := list[0]
        if first != head {
            prev.Next = first
            prev = first
        }
        list = list[1:]
        //fmt.Println(list)
        if len(list) > 0 {
            last := list[len(list)-1]
            list = list[:len(list)-1]
            prev.Next = last
            prev = last
            //fmt.Println(list)
        }
    }
    prev.Next = nil
}

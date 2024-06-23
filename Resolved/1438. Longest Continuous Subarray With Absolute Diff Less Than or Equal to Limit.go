// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/

import "container/heap"

type Item struct {
    Index int
    Value int
}


type MaxHeap []Item

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].Value > h[j].Value
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h MaxHeap) Peek() Item {
	if len(h) == 0 {
		return Item{}
	}
	return h[0]
}


// Min

type MinHeap []Item

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Value < h[j].Value
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h MinHeap) Peek() Item {
	if len(h) == 0 {
		return Item{}
	}
	return h[0]
}

func longestSubarray(nums []int, limit int) int {
    // use min and max heap to track min and max values and according indexes
    // we don't remove from heap, just start not considering elements from heap that has index outside of sliding window
    minH := MinHeap([]Item{})
    heap.Init(&minH)
    
    maxH := MaxHeap([]Item{})
	heap.Init(&maxH)

    l := 0

    result := 0

    for r := 0; r < len(nums); r++ {
        heap.Push(&minH, Item{Index: r, Value: nums[r]})
        heap.Push(&maxH, Item{Index: r, Value: nums[r]})

        for maxH.Peek().Value - minH.Peek().Value > limit {
            l++

            // skip elements with indicies less than l
            for len(minH) > 0 && minH.Peek().Index < l {
                heap.Pop(&minH)
            }
            for len(maxH) > 0 && maxH.Peek().Index < l {
                heap.Pop(&maxH)
            }
        }

        result = max(result, r - l + 1)
    }

    return result
}

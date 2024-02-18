// https://leetcode.com/problems/meeting-rooms-iii/

import "container/heap"

func mostBooked(n int, meetings [][]int) int {
    sort.Slice(meetings, func(i, j int) bool {
        return meetings[i][0] < meetings[j][0]
    })

    //fmt.Println(meetings)

    allRooms := []int{}
    for i := 0; i < n; i++ {
        allRooms = append(allRooms, i)
    }

    freeRooms := MinHeap(allRooms)
    heap.Init(&freeRooms)

    currMeetings := MinHeapMeeting([]Meeting{})
    heap.Init(&currMeetings)

    timesOccupied := make([]int, n)

    for _, m := range meetings {
        for currMeetings.Len() > 0 && currMeetings.Peek().End <= m[0] {
            prevM := heap.Pop(&currMeetings).(Meeting)
            heap.Push(&freeRooms, prevM.Room)
        }
        //fmt.Println("Meeting ", m)
        if freeRooms.Len() > 0 {
            r :=  heap.Pop(&freeRooms).(int)
            mt := Meeting{Start: m[0], End: m[1], Room: r}
            //fmt.Println("=>Free Room ", r)
            heap.Push(&currMeetings, mt)
            timesOccupied[r]++
        } else {            
            //fmt.Println("Current meetings: ", currMeetings)
            earliestM := heap.Pop(&currMeetings).(Meeting)
            r := earliestM.Room
            mt := Meeting{Start: earliestM.End, End: m[1] - m[0] + earliestM.End, Room: r}
            //fmt.Println("=>Occupies room of meeting ", earliestM, r)
            heap.Push(&currMeetings, mt)
            timesOccupied[r]++
        }
    }

    maxTime := 0
    maxRoom := 0

    fmt.Println(timesOccupied)

    for i, v := range timesOccupied {
        if v > maxTime {
            maxTime = v
            maxRoom = i
        }
    }

    return maxRoom
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h MinHeap) Peek() int {
	if len(h) == 0 {
		return -1
	}
	return h[0]
}

type Meeting struct {
    Start int
    End int
    Room int
}

type MinHeapMeeting []Meeting

func (h MinHeapMeeting) Len() int {
	return len(h)
}

func (h MinHeapMeeting) Less(i, j int) bool {
    if h[i].End == h[j].End {
        return h[i].Room < h[j].Room
    }
	return h[i].End < h[j].End
}
func (h MinHeapMeeting) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeapMeeting) Push(x interface{}) {
	*h = append(*h, x.(Meeting))
}
func (h *MinHeapMeeting) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h MinHeapMeeting) Peek() Meeting {
	if len(h) == 0 {
		return Meeting{}
	}
	return h[0]
}

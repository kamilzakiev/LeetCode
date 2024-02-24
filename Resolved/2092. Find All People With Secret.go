func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
    secret := make([]bool, n)

    meetings = append(meetings, []int{0, firstPerson, 0})
    secret[0] = true

    sort.Slice(meetings, func(i, j int) bool {
        return meetings[i][2] < meetings[j][2]
    })

    var prevMeeting []int
    sameTimeMeetings := [][]int{}
    for i, m := range meetings {
        if i == 0 {
            sameTimeMeetings = [][]int{m}
        } else if prevMeeting[2] != m[2] {
            processSameTime(n, sameTimeMeetings, secret)
            sameTimeMeetings = [][]int{m}
        } else {
            sameTimeMeetings = append(sameTimeMeetings, m)
        }
        if len(meetings) > 1 && i == len(meetings)-1 { // process last one
            processSameTime(n, sameTimeMeetings, secret)
        }
        prevMeeting = m
    }
    result := []int{}
    for i, b := range secret {
        if b {
            result = append(result, i)
        }
    }
    return result
}

func processSameTime(n int, meetings [][]int, secret []bool) {
    visited := make([]bool, n)
    queue := []int{}
    links := make(map[int][]int)

    for _, m := range meetings {
        from := m[0]

        if !visited[from] && secret[from] {
            queue = append(queue, from)
            visited[from] = true
        }

        from = m[1]

        if !visited[from] && secret[from] {
            queue = append(queue, from)
            visited[from] = true
        }
        

        if links[m[0]] == nil {
            links[m[0]] = []int{}
        }
        if links[m[1]] == nil {
            links[m[1]] = []int{}
        }
        links[m[0]]  = append(links[m[0]] , m[1])
        links[m[1]]  = append(links[m[1]] , m[0])
    }

    //fmt.Println(queue)
    //fmt.Println(links)

    for len(queue) > 0 {
        popped := queue[0]
        queue = queue[1:]

        for _, next := range links[popped] {
            if !visited[next] && !secret[next] {
                //fmt.Println(popped, " => ", next)
                secret[next] = true
                visited[next] = true
                queue = append(queue, next)
            }
        }
    }
}

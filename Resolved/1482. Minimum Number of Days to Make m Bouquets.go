// https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/
func minDays(bloomDay []int, m int, k int) int {
    n := len(bloomDay)
    if int64(n) < int64(m) * int64(k) {
        return -1
    }
    max := 0
    for _, v := range bloomDay {
        if max < v {
            max = v
        }
    }
    // binary search
    l, r := 0, max

    for l <= r {
        mid := (l+r)/2
        if canGatherFlowers(mid, bloomDay, m, k) {
            //fmt.Println(mid, true)
            r = mid - 1
        } else {
            //fmt.Println(mid, false)
            l = mid + 1
        }
    }
    //fmt.Println(l, r, max)
    if l > max {
        return -1
    }
    return l
}

// calculate if we can gather flowers on specific day
func canGatherFlowers(day int, bloomDay []int, m int, k int) bool {
    consecCount := 0
    bouquetCount := 0
    for _, v := range bloomDay {
        if v <= day {
            consecCount++
        } else {
            consecCount = 0
        }
        if consecCount == k {
            consecCount = 0
            bouquetCount++
        }
        if bouquetCount >= m {
            return true
        }
    }
    //fmt.Println(consecCount, bouquetCount, m, k)
    return false
}

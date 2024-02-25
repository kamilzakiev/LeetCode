// https://leetcode.com/problems/ugly-number-ii/
func nthUglyNumber(n int) int {
    list := []int{1}
    p2, p3, p5 := 0, 0, 0 // pointers for 2, 3 and 5 multiplier within list
    // move one of 3 pointer if it produces the smallest number among 3
    // add it to the list only if list does not have this number already
    
    for len(list) < n {
        m2 := list[p2]*2
        m3 := list[p3]*3
        m5 := list[p5]*5

        if m2 <= min(m3, m5) {
            if list[len(list)-1] < m2 {
                list = append(list, m2)
            }
            p2++
        } else if m3 <= min(m2, m5) {
            if list[len(list)-1] < m3 {
                list = append(list, m3)
            }
            p3++
        } else {
            if list[len(list)-1] < m5 {
                list = append(list, m5)
            }
            p5++
        }
    }
    return list[n-1]
}

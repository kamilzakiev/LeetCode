// https://leetcode.com/problems/super-ugly-number/
func nthSuperUglyNumber(n int, primes []int) int {
    // keep index of each prime within final list of super ugly numbers
    // take the smallest val and index++ for that prime
    pN := len(primes)
    indicies := make([]int, pN)
    list := []int{1}
    for len(list) < n {
        minI := 0
        minMult := math.MaxInt32
        for i, v := range primes {
            index := indicies[i]
            mult := v * list[index]
            if mult < minMult {
                minMult = mult
                minI = i
            }
        }
        indicies[minI]++
        if list[len(list)-1] != minMult {
            list = append(list, minMult)
        }
    }
    //fmt.Println(list)
    return list[n-1]
}

// https://leetcode.com/problems/largest-divisible-subset/
func largestDivisibleSubset(nums []int) []int {
    sort.Ints(nums)
    dividends := make(map[int][]int)

    for i := 0; i < len(nums)-1; i++ {
        dividends[nums[i]] = []int{}
        for j := i+1; j < len(nums); j++ {
            if nums[j] % nums[i] == 0 {
                dividends[nums[i]] = append(dividends[nums[i]], nums[j])
            }
        }
    }
    //fmt.Println(dividends)
    result := []int{nums[0]}

    maxDividendsMemo := make(map[int][]int)
    
    var iter func(v int) []int
    iter = func(v int) []int {
        if prevValue, exists := maxDividendsMemo[v]; exists {
            return prevValue
        }

        divResult := []int{v}
        for _, dvd := range dividends[v] {
            parentDividends := iter(dvd)
            if len(parentDividends) + 1 > len(divResult) {
                divResult = append([]int{v}, parentDividends...)
            }
        }

        if len(divResult) > len(result) {
            result = divResult
        }

        maxDividendsMemo[v] = divResult
        return divResult
    }

    for _, v := range nums {
        iter(v)
    }
    //fmt.Println(dividendsMemo)

    return result
}

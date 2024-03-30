// https://leetcode.com/problems/subarrays-with-k-different-integers/
func subarraysWithKDistinct(nums []int, k int) int {
    // tip: result is number of subarrays with at most k distinct - number of subarrays with at most k-1 distinct
    return int(AtMostKDistinct(nums, k) - AtMostKDistinct(nums, k-1))
}

func AtMostKDistinct(nums []int, k int) int64 {        
        result := int64(0)
        start := 0
        distCountInWindow := 0
        mapI := make(map[int]int)

        for end := 0; end < len(nums); end++ {
            if mapI[nums[end]] == 0 {
                distCountInWindow++
            }
            mapI[nums[end]]++

            for k+1 == distCountInWindow {
                mapI[nums[start]]--
                if mapI[nums[start]] == 0 {
                    distCountInWindow--
                }
                start++
            }
            fmt.Println(start, end, distCountInWindow)
            result += int64(end - start)
        }

        return result    
}

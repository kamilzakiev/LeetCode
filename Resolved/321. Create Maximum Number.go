// https://leetcode.com/problems/create-maximum-number/
func maxNumber(nums1 []int, nums2 []int, k int) []int {
    // use monothonic stack, but not simple, we need that one that provides k numbers
    // for all possible length of monothonic stacks for nums1 and nums2, we need len(stack1) + len(stack2) = k
    // then combine these stacks
    maxArray := []int{}
    for len1 := 0; len1 <= len(nums1); len1++ {
        len2 := k - len1
        if len2 <= len(nums2) && len2 >= 0 {
            stack1 := myMaxNumber(nums1, len1)
            stack2 := myMaxNumber(nums2, len2)
            
            combined := maxCombine(stack1, stack2)

            //fmt.Println(stack1, stack2, combined)

            maxArray = getBigger(maxArray, combined)
        }
    }

    return maxArray
}

func getBigger(currMax, candidate []int) []int {
    if len(currMax) == 0 {
        return candidate
    }
    // they are the same length
    for i := range currMax {
        if currMax[i] < candidate[i] {
            return candidate
        } else if currMax[i] > candidate[i] {
            return currMax
        }
    }
    return currMax

}

func maxCombine(nums1, nums2 []int) []int {
    i1, i2 := 0, 0

    result := []int{}
    for i1 < len(nums1) && i2 < len(nums2) {
        choose1, choose2 := false, false        
        if nums1[i1] == nums2[i2] {
            j1, j2 := i1, i2
            for j1 < len(nums1) && j2 < len(nums2) {
                if nums1[j1] > nums2[j2] {
                    choose1 = true
                    break
                } else if nums1[j1] < nums2[j2] {
                    choose2 = true
                    break
                }
                j1++
                j2++
            }
            if !choose1 && !choose2 {
                if j1 < len(nums1) {
                    choose1 = true
                } else {
                    choose2 = true
                }
            }
        } else if nums1[i1] > nums2[i2] {
            choose1 = true            
        } else {
            choose2 = true            
        }
        if choose1 {
            result = append(result, nums1[i1])
            i1++
        } else {
            result = append(result, nums2[i2])
            i2++
        }
    }

    for i1 < len(nums1) {
        result = append(result, nums1[i1])
        i1++
    }

    for i2 < len(nums2) {
        result = append(result, nums2[i2])
        i2++
    }

    return result
}

// returns max sequence of k length (if len(nums) >= k)
// monotonic stack but also considers k
func myMaxNumber(nums []int, k int) []int {
    if k == 0 {
        return []int{}
    }
    n := len(nums)
    result := []int{}

    for i, v := range nums {
        for len(result) > 0 && result[len(result)-1] < v && len(result) + (n-i-1) >= k {
            result = result[:len(result)-1]
        }
        if len(result) < k {
            result = append(result, v)
        }
    }

    return result
}

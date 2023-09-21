//https://leetcode.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    fmt.Println("---------------------")
    fmt.Println(nums1, nums2)
    var medNum = (len(nums1) + len(nums2)) / 2

    if ((len(nums1) + len(nums2)) % 2 == 0) {
        medNum = medNum-1
    }
    var result = float64(getNthElement(medNum, nums1, nums2))
    
    fmt.Printf("medNum=%d, result1=%f\r\n", medNum, result)

    if ((len(nums1) + len(nums2)) % 2 == 0) {
        medNum = medNum+1
        result2 := float64(getNthElement(medNum, nums1, nums2))
        fmt.Printf("medNum=%d, result2=%f\r\n", medNum, result2)
        result = float64((result+result2)/2)
    }
    return result
}

func getNthElement(medNum int, nums1 []int, nums2 []int) int {
    // zero length corner cases
    if len(nums1) == 0 {
        return nums2[medNum]
    }
    if len(nums2) == 0 {
        return nums1[medNum]
    }
    // corner cases, when medNum is not in intersection of nums1 and nums2
    if medNum < len(nums1) && nums1[medNum] <= nums2[0] {return nums1[medNum]}
    if medNum < len(nums2) && nums2[medNum] <= nums1[0] {return nums2[medNum]}
    if medNum >= len(nums1) && nums2[medNum-len(nums1)] > nums1[len(nums1)-1] {return nums2[medNum-len(nums1)]}
    if medNum >= len(nums2) && nums1[medNum-len(nums2)] > nums2[len(nums2)-1] {return nums1[medNum-len(nums2)]}

    fmt.Println("--------First---------")
    nums1Index, found := findElementIndex(0, len(nums1)-1, medNum, nums1, nums2)

    if found {
        return nums1[nums1Index]
    } else {
        fmt.Println("--------Second---------")
        nums2Index, _ := findElementIndex(0, len(nums2)-1, medNum, nums2, nums1)
        return nums2[nums2Index]
    }
}

func findElementIndex(left int, right int, medNum int, nums1 []int, nums2 []int) (int, bool) {
    fmt.Printf("medNum=%d\r\n", medNum)
    fmt.Printf("left/right=(%d,%d)\r\n", left, right)
    

    midNums1Index := (left + right)/2
    midNums2Index := min(max(0, medNum - midNums1Index-1), len(nums2)-1)
    //midNums1Index = medNum - midNums2Index - 1
    fmt.Printf("Nums1Index/Nums2Index=(%d,%d)\r\n", midNums1Index, midNums2Index)

    isCorrectMasterNum := midNums1Index == medNum - midNums2Index - 1 &&  isCorrectMedNum(midNums1Index, midNums2Index, nums1, nums2)
    fmt.Printf("isCorrectMasterNum=%t\r\n", isCorrectMasterNum)

    if left >= right {
        return midNums1Index, isCorrectMasterNum
    }
    if isCorrectMasterNum {
        return midNums1Index, true
    }
    if nums1[midNums1Index] < nums2[midNums2Index] {
        return findElementIndex(midNums1Index+1, right, medNum, nums1, nums2)
    } else {
        return findElementIndex(left, midNums1Index-1, medNum, nums1, nums2)
    }
}

func isCorrectMedNum(masterIndex int, slaveIndex int, master []int, slave []int) bool {
    if slaveIndex < 0 || slaveIndex > len(slave) - 1 {
        return false
    }
    if slave[slaveIndex] > master[masterIndex] {
        return false
    }
    if slaveIndex <  len(slave) - 1 && slave[slaveIndex + 1] < master[masterIndex] {
        return false
    }
    return true
}


func min(a, b int) int {
    if (a < b) {
        return a
    }
    return b
}

func max(a, b int) int {
    if (a < b) {
        return b
    }
    return a
}

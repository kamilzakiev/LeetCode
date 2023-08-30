// https://leetcode.com/problems/minimum-replacements-to-sort-the-array
func minimumReplacement(nums []int) int64 {
    // just go from tail to start and split number based on previous value
    // ensure that after split min element among new numbers will be max as possible
    n := len(nums)
    prev := nums[n-1]

    result := int64(0)
    for i := n-2; i >=0; i-- {
        v := nums[i]

        if v > prev {
            div := v / prev // how many numbers will be produced
            rem := v % prev
            if rem > 0 {
                div++
            }
            // once we know number of produced numbers, our goal is to make min value of them maximum
            // for example, if we split 7, it would be better to split it as 3+4 instead of 1+6
            // let's assume we need to split 20 to 3 elements, then they should be 6, 7, 7
            // 6 will be prev, and let's call 7 as maxNewElement
            prev = v / div // based on hard math calculations            

            //fmt.Println(fmt.Sprintf("%d is split to starting value %d and total %d numbers", v, prev, div))

            result += int64(div-1)
        } else {
            prev = v
           // fmt.Println(fmt.Sprintf("%d not split", v))
        }
    }

    return result
}

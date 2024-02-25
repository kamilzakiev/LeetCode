// https://leetcode.com/problems/single-number-iii
func singleNumber(nums []int) []int {
    xor := 0
    for _, v := range nums {
        xor = xor ^ v
    }
    diffIndex := 0
    for i := 0; i < 32; i++ {
        if xor & (1 << i) > 0 {
            diffIndex = i
            break
        }
    }

    xor0 := 0
    xor1 := 0
    for _, v := range nums {
        if v & (1 << diffIndex) == 0 {
            xor0 = xor0 ^ v
        } else {
            xor1 = xor1 ^ v
        }
    }

    return []int{xor0, xor1}
}

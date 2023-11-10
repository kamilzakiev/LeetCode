// https://leetcode.com/problems/decoded-string-at-index/
func decodeAtIndex(s string, k int) string {
    n := len(s)
    length := make([]int64, n)

    currLength := int64(0)
    for i, r := range s {
        if r >= 'a' && r <= 'z' {
            currLength++
        } else {
            m := int(r - '0')
            currLength = currLength * int64(m)
        }
        length[i] = currLength

        if currLength == int64(k) {
            return findLastRuneOccurence(s, i)
        }
        if currLength > int64(k) {
            return findKthRune(s, k, i, length)
        }
    }

    return ""
}

// find kth rune, knowing that by maxIndex we have string that is longer than k, 
// but maxIndex-1 provides shorter string than k
func findKthRune(s string, k int, maxIndex int, length []int64) string {
    prevValue := length[maxIndex-1] // i cannot be 0 here, because k >= 1
    // we know that prevValue < k
    k = int(int64(k) % prevValue)
    if k == 0 {
        return findLastRuneOccurence(s, maxIndex-1)
    }
    for length[maxIndex-1] > int64(k) {
        maxIndex--
    }
    if length[maxIndex-1] == int64(k) {
        return findLastRuneOccurence(s, maxIndex-1)
    }
    return findKthRune(s, k, maxIndex, length)
}

// returns closest rune to i as string
func findLastRuneOccurence(s string, i int) string {
    for s[i] < 'a' || s[i] > 'z' {
        i--
    }
    return string(s[i])
}

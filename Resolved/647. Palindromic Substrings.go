// https://leetcode.com/problems/palindromic-substrings/
func countSubstrings(s string) int {
    // for every character, expand simultaneously to the left and right
    // that will give us O(n^2)
    result := 0
    for i := range s { // for every character, expand to righ and left to get palindrome
        result++
        for r,l := i+1, i-1; l >= 0 && r < len(s); { // this one for odd length
            if s[l] == s[r] {
                result ++
                l--
                r++
            } else {
                break
            }
        }
    }
    for i := range s { // for every character, expand to righ and left to get palindrome
        for r,l := i+1, i; l >= 0 && r < len(s); { // this one for even length
            if s[l] == s[r] {
                result ++
                l--
                r++
            } else {
                break
            }
        }
    }
    return result
}

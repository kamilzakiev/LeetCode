// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
func minRemoveToMakeValid(s string) string {
     // My own algorithm:
    // 1. Add all '(' to result
    // 2. Add ')' only if number of '(' is more than ')'
    // 3. Once completed, if number of '(' is more than number of ')', then remove extensive '(' from the end   
    openCount := 0

    resultB := []byte(s)

    removedCount := 0

    for i := 0; i < len(resultB); i++ {
        b := resultB[i]
        if b == '(' {
            openCount++
        } else if b == ')' {
            if openCount > 0 {
                openCount--
            } else {
                resultB[i] = 0
                removedCount++
            }
        }
    }

    if openCount > 0 {
        for i := len(resultB)-1; i >= 0; i-- {
            if resultB[i] == '(' && openCount > 0 {
                resultB[i] = 0
                removedCount++
                openCount--
            }
        }
    }

    resultList := make([]byte, len(s) - removedCount)
    j := 0

    for i := 0; i < len(resultB); i++ {
        b := resultB[i]
        if b != 0 {
            resultList[j] = b
            j++
        }
    }
    return string(resultList)
}

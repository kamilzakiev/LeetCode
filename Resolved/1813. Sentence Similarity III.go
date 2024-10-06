// https://leetcode.com/problems/sentence-similarity-iii/
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
    if sentence1 == sentence2 {
        return true
    }
    sh, lg := sentence1, sentence2
    if len(sentence1) > len(sentence2) {
        sh, lg = sentence2, sentence1
    }
    wSh := strings.Split(sh, " ")
    wLg := strings.Split(lg, " ")

    // 3 cases:
    //"Eating right now"
    //"Eating"

    //"Eating right now"
    //"Eating now"

    //"Eating right now"
    //"right now"

    // to be similar, they should have common prefix and suffix. and short should consist of only prefix + suffix

    prefixWordCount := 0
    for i := range wSh {
        if wSh[i] == wLg[i] {
            prefixWordCount++
        } else {
            break
        }
    }
    suffixWordCount := 0
    j := len(wLg) - 1 // index for long one
    for i := len(wSh) - 1; i >= 0; i-- {
        if wSh[i] == wLg[j] {
            suffixWordCount++
            j--
        } else {
            break
        }
    }
    return prefixWordCount + suffixWordCount >= len(wSh) // > for the case: "a aa a" and "a"
}

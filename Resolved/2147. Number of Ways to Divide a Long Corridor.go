// https://leetcode.com/problems/number-of-ways-to-divide-a-long-corridor/
func numberOfWays(corridor string) int {
    // idea - dividers can be set only between pairs of seats
    // so for each space between 2 pairs, count how many ways to install 1 divider
    const mod = 1e9+7
    seatCount := 0
    result := 1
    for i := 0; i < len(corridor); i++ {
        b := corridor[i]
        if b == 'S' {
            seatCount++
            if seatCount % 2 == 0 {
                // one pair of seats just ended here
                // count how many plants
                plantCount := 0
                i++
                for i < len(corridor) && corridor[i] == 'P' {
                    plantCount++
                    i++
                }
                i--
                //fmt.Println(i, plantCount)
                if i < len(corridor)-1 {
                    result = int((int64(result) * int64(plantCount + 1)) % mod) 
                }
            }
        }
    }
    if seatCount < 2 || seatCount % 2 != 0 {
        return 0
    }
    return result
}

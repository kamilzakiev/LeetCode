// https://leetcode.com/problems/fraction-to-recurring-decimal/
func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {
        return "0"
    }
    negative := numerator < 0 && denominator > 0 || numerator > 0 && denominator < 0
    numerator = abs(numerator)
    denominator = abs(denominator)

    whole := numerator / denominator        

    result := strconv.Itoa(whole)

    numerator -= whole * denominator

    if numerator > 0 {
        result += "."
        end := false
        mapNumerators := make(map[int]int)
        for !end { 
            if mapNumerators[numerator] > 0 {
                i := mapNumerators[numerator]
                result = result[:i] + "(" + result[i:]
                result += ")"
                end = true
            } else {
                mapNumerators[numerator] = len(result)            
                numerator *= 10
                for numerator < denominator {
                    numerator *= 10
                    result += "0"
                }            
                whole = numerator / denominator
                result += strconv.Itoa(whole)
                numerator -= whole * denominator
                if numerator == 0 {
                    end = true
                } 
            }
        }
    }

    if negative {
        result = "-" + result
    }
    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

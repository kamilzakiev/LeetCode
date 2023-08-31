// https://leetcode.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/
func minTaps(n int, ranges []int) int {
   // for each point find a tap that covers most number of points to the right
   // start from 0
   /*sort.Slice(ranges, func(i, j int) bool {
       return ranges[i] + i < ranges[j] + j
   }) */

   rangeEnds := map[int]int{} // value - range end, key - index of the longest range ending at value

   for i, v := range ranges {
       rangeEnd := i + v
       if prevValue, exists := rangeEnds[rangeEnd]; exists {
           if prevValue > i { // need to use range with the most topleft center
               ranges[rangeEnd] = i 
           }
       } else {
           rangeEnds[rangeEnd] = i 
       }
   }  

   //fmt.Println(rangeEnds)

   rangeEnd := -1
   result := 0
   for point := 0; point < n; point++ { // ensure interval [point, point+1] is covered 
       //fmt.Println(point, rangeEnd)
       if point + 1 > rangeEnd {
           // find tap with most right range
           found := false
            for i := point + 200; i > point; i-- { //200, because 0 <= ranges[i] <= 100. i==0 is not considered, since point+1 also should be covered
                if rangeIndex, exists := rangeEnds[i]; exists {
                    rangeStart := rangeIndex - ranges[rangeIndex]
                    if rangeStart <= point {
                        rangeEnd = ranges[rangeIndex] + rangeIndex // or =i
                        //fmt.Println("Tap ", point, "is covered by", rangeIndex, ":[", rangeStart, ":", rangeEnd, "]")
                        result++
                        found = true
                        break
                    }
                }
            }
            if !found {
                return -1
            }
       }
   }

   return result
}

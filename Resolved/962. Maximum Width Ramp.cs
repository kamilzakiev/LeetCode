// https://leetcode.com/problems/maximum-width-ramp/
public class Solution {
    public int MaxWidthRamp(int[] nums) {
        // 1. Sort nums by values while keeping indicies
        // 2. Find max diff between min and max indices withing sorted list
        var sorted = nums.ToArray().Select((v, i) => new Tuple<int, int>(v, i)).OrderBy(t => t.Item1);
        var enumerator = sorted.GetEnumerator();
        
        enumerator.MoveNext();
        var minSoFarIndex = enumerator.Current.Item2;

        var maxDelta = 0;
        while(enumerator.MoveNext())
        {
            var currIndex = enumerator.Current.Item2;
            if (currIndex < minSoFarIndex)
            {
                minSoFarIndex = currIndex;
            }
            maxDelta = Math.Max(maxDelta, currIndex - minSoFarIndex);
        }
        return maxDelta;
    }
}

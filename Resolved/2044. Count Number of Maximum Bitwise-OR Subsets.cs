// https://leetcode.com/problems/count-number-of-maximum-bitwise-or-subsets/
public class Solution {
    public int CountMaxOrSubsets(int[] nums) {
        var n = nums.Length;
        var orValue = 0;

        foreach(var i in nums)
        {
            orValue |= i;
        }
        var result = 0; 
        void Backtrack(int i, int currentOrValue) {
            if (i == n) {
                return;
            }
            // without current element
            Backtrack(i + 1, currentOrValue);
            // with current element
            currentOrValue |= nums[i];
            if (currentOrValue == orValue)
            {
                result++;
            }
            Backtrack(i + 1, currentOrValue);
        };

        Backtrack(0, 0);

        return result;
    }
}

// https://leetcode.com/problems/longest-square-streak-in-an-array/
public class Solution {
    public int LongestSquareStreak(int[] nums) {
        var n = nums.Length;
        Array.Sort(nums);
        var parentIndices = new Dictionary<int, int>();
        for (var i = 0; i < n; i++)
        {
            parentIndices[nums[i]] = i;
        }
        for (var i = 0; i < n; i++)
        {
            var val = nums[i];
            var sq = GetIntSquareRoot(val);
            if (sq.HasValue)
            {
                if (parentIndices.TryGetValue(sq.Value, out int parentIndex))
                {
                    parentIndices[val] = parentIndex;
                }
            }
        }
        var children = new Dictionary<int, int>(); // number of children for each parentIndex
        var result = 0;
        foreach (var pair in parentIndices)
        {
            if (children.TryGetValue(pair.Value, out int preVal))
            {
                children[pair.Value] = preVal + 1;
            } 
            else 
            {
                children[pair.Value] = 1;
            }
            result = Math.Max(result, children[pair.Value]);
        }
        if (result == 1)
        {
            return -1;
        }
        return result;
    }

    private static int? GetIntSquareRoot(int a) 
    {
        int t = (int)Math.Sqrt(a);
        if (t * t == a)
        {
            return t;
        }
        return null;
    }
}

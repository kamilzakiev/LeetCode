// https://leetcode.com/problems/minimum-number-of-removals-to-make-mountain-array/
public class Solution {
    public int MinimumMountainRemovals(int[] nums) {
        var n = nums.Length;

        var result = 0;
        for (var i = 1; i < n-1; i++)
        {
            var lisLength = LongestIncSubsequence(new Span<int>(nums, 0, i + 1));
            var ldsLength = LongestDecSubsequence(new Span<int>(nums, i, n - i));

            if (lisLength > 1 && ldsLength > 1) 
            {
                result = Math.Max(result, lisLength + ldsLength - 1);
            }
            //Console.WriteLine($"{i}: {lisLength}, {ldsLength}");

        }
        return n - result;
    }

    // see https://leetcode.com/problems/longest-increasing-subsequence/
    private static int LongestIncSubsequence(Span<int> nums)
    {
        return LongestSubsequence(nums, Comparer<int>.Default);
    }

    private static int LongestDecSubsequence(Span<int> nums)
    {
        return LongestSubsequence(nums, Comparer<int>.Create((x, y) => y.CompareTo(x)));
    }
    
    private static int LongestSubsequence(Span<int> nums, IComparer<int> comparer)
    {
        int n = nums.Length;
        List<int> ans = new List<int>();

        ans.Add(nums[0]);

        for (int i = 1; i < n; i++)
        {
            if (comparer.Compare(nums[i], ans[ans.Count - 1]) > 0)
            {
                ans.Add(nums[i]);
            }
            else
            {
                int pos = ans.BinarySearch(nums[i], comparer);
                if (pos < 0)
                {
                    pos = ~pos;
                }
                ans[pos] = nums[i];
            }
        }
        return ans.Count;
    }   
}

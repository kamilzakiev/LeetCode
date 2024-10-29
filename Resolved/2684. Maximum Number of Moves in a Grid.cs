// https://leetcode.com/problems/maximum-number-of-moves-in-a-grid/
public class Solution {
    public int MaxMoves(int[][] grid) {
        int m = grid.Length, n = grid[0].Length;
        var dp = new int[m, n];

        var result = 0; 

        for (var j = n-2; j >= 0; j--)
        {
            for (var i = 0; i < m; i++)
            {
                var val = grid[i][j];
                var max = 0;
                if (i > 0 && grid[i-1][j+1] > val)
                {
                    max = dp[i-1, j+1] + 1;
                }
                if (grid[i][j+1] > val && dp[i, j+1] + 1 > max)
                {
                    max = dp[i, j+1] + 1;
                }
                if (i < m-1 && grid[i+1][j+1] > val && dp[i+1, j+1] + 1 > max)
                {
                    max = dp[i+1, j+1] + 1;
                }
                dp[i, j] = max;
                if (j == 0)
                {
                    result = Math.Max(result, dp[i, j]);
                }
            }
        }

        return result;
    }
}

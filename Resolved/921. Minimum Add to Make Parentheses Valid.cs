// https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/
public class Solution {
    public int MinAddToMakeValid(string s) {
        var nonValidCloseBracketCount = 0;
        var totalSum = 0;
        foreach(var c in s)
        {
            if (c == '(') 
            {
                totalSum++;
            } else {
                if (totalSum <= 0) {
                    nonValidCloseBracketCount++;
                } else {
                    totalSum--;
                }
            }
        }
        // we need to add opening for each non valid closing bracket, and close every non-closed opening bracket
        return nonValidCloseBracketCount + Math.Abs(totalSum); 
    }
}

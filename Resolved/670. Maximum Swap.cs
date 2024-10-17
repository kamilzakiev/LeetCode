// https://leetcode.com/problems/maximum-swap/
public class Solution {
    public int MaximumSwap(int num) {
        var digits = num.ToString().ToCharArray().Select(d => (int)d - (int)'0').ToList();
        var maxIndex = new Dictionary<int, int>(); // max index for each digit 0..9
        for (var i = 0; i < digits.Count; i++) 
        {
            var digit = digits[i];
            maxIndex[digit] = i;
        }

        for (var i = 0; i < digits.Count; i++) 
        {
            var digit = digits[i];
            var found = false;
            for (var maxDigit = 9; maxDigit > digit; maxDigit--)
            {
                if (maxIndex.TryGetValue(maxDigit, out var maxIndexValue) && maxIndexValue > i)
                {
                    // here we found a swap
                    var temp = digits[i];
                    digits[i] = maxDigit;
                    digits[maxIndexValue] = temp;
                    found = true;
                    break;
                }
            }
            if (found)
            {
                break;
            }
        }
        var result = 0;
        for (var i = 0; i < digits.Count; i++) 
        {
            result = result * 10 + digits[i];
        } 
        return result;
    }
}

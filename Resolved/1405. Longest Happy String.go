// https://leetcode.com/problems/longest-happy-string/
public class Solution {
    public string LongestDiverseString(int a, int b, int c) {
        var dict = new Dictionary<char, int>();
        dict['a'] = a;
        dict['b'] = b;
        dict['c'] = c;

        char prevChar = '0';
        var exit = false;
        var result = new StringBuilder();

        while(!exit)
        {
            var maxChar = GetMaxChar(dict, prevChar);
            if (!maxChar.HasValue)
            {
                exit = true;
                break;
            }

            if (dict[maxChar.Value] == 1 
                || prevChar != '0' && dict[prevChar] > dict[maxChar.Value])
            {
                result.Append(maxChar.Value);
                dict[maxChar.Value]--;
            } 
            else
            {
                result.Append(maxChar.Value);
                result.Append(maxChar.Value);

                dict[maxChar.Value] = dict[maxChar.Value] - 2;
            }

            prevChar = maxChar.Value;
        }
        return result.ToString();
        
    }

    private static char? GetMaxChar(Dictionary<char, int> dict, char exceptChar)
    {
        char maxChar = '0';
        var maxCharCount = 0;

        foreach (var pair in dict)
        {
            if (pair.Key != exceptChar && pair.Value > maxCharCount)
            {
                maxChar = pair.Key;
                maxCharCount = pair.Value;
            }
        }
        if (maxCharCount == 0)
        {
            return null;
        }
        return maxChar;
    }
}

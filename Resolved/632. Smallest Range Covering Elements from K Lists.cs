// https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/

public class Solution {
    private class ValueIfList 
    {
        public int Value  {get; set;}
        public int ListNum {get; set;}
        public int Index {get; set;}
    }
    public int[] SmallestRange(IList<IList<int>> nums) {
        // use sliding window with elements from all lists
        // use priority queue to track min element
        var n = nums.Count;
        var valueIndexInList = new int[n];

        var pq = new PriorityQueue<ValueIfList, int>();

        var maxValue = Int32.MinValue; 
        for(var i = 0; i < n; i++)
        {
            pq.Enqueue(new ValueIfList(){Value = nums[i][0], ListNum = i, Index = 0}, nums[i][0]);
            if (maxValue < nums[i][0]) 
            {
                maxValue = nums[i][0];
            }
        }
        var minValue = pq.Peek().Value;
        var resultLength = maxValue - minValue;
        var result = new int[]{minValue, maxValue};

        // while list of the min element is not processed fully

        while (pq.Peek().Index < nums[pq.Peek().ListNum].Count - 1)
        {
            var removed = pq.Dequeue();
            var removedListNum = removed.ListNum;
            var removedItemIndex = removed.Index;
            var itemToAdd = nums[removedListNum][removedItemIndex + 1];
            pq.Enqueue(new ValueIfList(){Value = itemToAdd, ListNum = removedListNum, Index = removedItemIndex + 1}, itemToAdd);
            if (maxValue < itemToAdd)
            {
                maxValue = itemToAdd;
            }
            minValue = pq.Peek().Value;
            if (resultLength > maxValue - minValue)
            {
                resultLength = maxValue - minValue;
                result[0] = minValue;
                result[1] = maxValue;
            }
        }

        return result;
    }
}

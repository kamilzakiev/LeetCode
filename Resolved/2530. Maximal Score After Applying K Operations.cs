// https://leetcode.com/problems/maximal-score-after-applying-k-operations/
public class Solution {
    public long MaxKelements(int[] nums, int k) {
        var comparer = Comparer<int>.Create((x, y) => y.CompareTo(x));
        var pq = new PriorityQueue<int, int>(comparer);

        foreach(var v in nums)
        {
            pq.Enqueue(v, v);
        }
        var result = 0L;
        for (var i = 0; i < k && pq.Count > 0; i++) {
            var max = pq.Dequeue();
            result += max;
            var newVal = (int)Math.Ceiling((decimal)max / 3.0M);
            pq.Enqueue(newVal, newVal);
        }
        return result;
    }
}

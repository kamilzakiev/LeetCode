// https://leetcode.com/problems/height-of-binary-tree-after-subtree-removal-queries/
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     public int val;
 *     public TreeNode left;
 *     public TreeNode right;
 *     public TreeNode(int val=0, TreeNode left=null, TreeNode right=null) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
public class Solution {
    public int[] TreeQueries(TreeNode root, int[] queries) {
        var nodeInfos = new Dictionary<int, NodeInfo>();
        // 1. Calculate NodeInfo.Height
        void CalcHeight(TreeNode node, int currHeight)
        {
            if (node == null)
            {
                return;
            }
            var nodeInfo = new NodeInfo{Height = currHeight};
            nodeInfos[node.val] = nodeInfo;
            CalcHeight(node.left, currHeight + 1);
            CalcHeight(node.right, currHeight + 1);
        }
        CalcHeight(root, 0);
        // 2. Calculate NodeInfo.SubtreeHeight
        int CalcSubtreeHeight(TreeNode node)
        {
            if (node == null)
            {
                return 0;
            }
            var subtreeHeight = Math.Max(nodeInfos[node.val].Height, Math.Max(CalcSubtreeHeight(node.left), CalcSubtreeHeight(node.right)));
            nodeInfos[node.val].SubtreeHeight = subtreeHeight;
            return subtreeHeight;
        }
        CalcSubtreeHeight(root);
        // 3. Calculate NodeInfo.OtherSubtreesMaxHeight
        void CalcOtherSubtreesMaxHeight(TreeNode parent, TreeNode node, bool isLeft)
        {
            if (node == null)
            {
                return;
            }
            var otherSubtreesMaxHeight = -1;
            if (parent != null)
            {
                otherSubtreesMaxHeight = nodeInfos[parent.val].OtherSubtreesMaxHeight;
                if (isLeft && parent.right != null)
                {
                    otherSubtreesMaxHeight = Math.Max(otherSubtreesMaxHeight, nodeInfos[parent.right.val].SubtreeHeight);
                }
                if (!isLeft && parent.left != null)
                {
                    otherSubtreesMaxHeight = Math.Max(otherSubtreesMaxHeight, nodeInfos[parent.left.val].SubtreeHeight);
                }
            }
            nodeInfos[node.val].OtherSubtreesMaxHeight = otherSubtreesMaxHeight;
            CalcOtherSubtreesMaxHeight(node, node.left, true);
            CalcOtherSubtreesMaxHeight(node, node.right, false);
        }
        CalcOtherSubtreesMaxHeight(null, root, false);
        // Now return query results
        var result = new int[queries.Length];
        for(var i = 0; i < queries.Length; i++)
        {
            var queryNode = queries[i];
            var nodeInfo = nodeInfos[queryNode];
            result[i] = Math.Max(nodeInfo.Height - 1, nodeInfo.OtherSubtreesMaxHeight);
        }
        return result;
    }

    private class NodeInfo
    {
        public int Height {get; set;}
        public int SubtreeHeight {get; set;}
        // max (parent.OtherSubtreeMaxHeight, neighb.SubtreeHeight), where neighb shares same parent. -1, if no other subtree
        public int OtherSubtreesMaxHeight {get; set;}
    }
}

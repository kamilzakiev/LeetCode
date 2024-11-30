// https://leetcode.com/problems/valid-arrangement-of-pairs/
public class Solution {
    public int[][] ValidArrangement(int[][] pairs) {
        // use Hierholzer’s algorithm to find Eulerian path
        var graph = new Graph();

        // find starting vertex
        // it has 1 more output that inputs

        var inputs = new Dictionary<int, int>();
        var outputs = new Dictionary<int, int>();

        foreach(var pair in pairs)
        {
            graph.AddLink(pair[0], pair[1]);
            inputs[pair[1]] = inputs.GetValueOrDefault(pair[1]) + 1;
            outputs[pair[0]] = outputs.GetValueOrDefault(pair[0]) + 1;
        }

        var start = -1;

        foreach(var link in graph.Children)
        {
            if (inputs.GetValueOrDefault(link.Key) + 1 == outputs.GetValueOrDefault(link.Key))
            {
                start = link.Key;
                break;
            }
        }

        if (start == -1)
        {
            foreach(var link in graph.Children)
            {
                start = link.Key;
                break;
            }
        }

        var path = graph.EulerianPath(start);

        var result = new List<int[]>();

        for (var i = 1; i < path.Count; i++)
        {
            result.Add(new int[2]{path[i-1], path[i]});
        }
        return result.ToArray();
    }
}

public class Graph
{
    public Graph()
    {
        this.Children = new Dictionary<int, List<(int, int)>>();
    }

    public Dictionary<int, List<(int, int)>> Children {get; set; }

    public void AddLink(int from, int to)
    {
        this.AddLink(from, to, 0);
    }

    public void AddLink(int from, int to, int dist)
    {
        if (!this.Children.TryGetValue(from, out List<(int, int)> children))
        {
            children = new List<(int, int)>();
            this.Children[from] = children;
        }
        children.Add((to, dist));
    }

    public int[] Dijkstra(int n, int src)
    {
        int[] dist = new int[n];
        for (int i = 0; i < n; i++) {
            dist[i] = int.MaxValue;
        }
        dist[src] = 0;

        // whether vertex is included into the shortest path tree or not
        var sptSet = new bool[n];

        // Find shortest path for all vertices
        for (int count = 0; count < n - 1; count++) {
            // Pick the minimum distance vertex 
            // from the set of vertices not yet
            // processed. u is always equal to
            // src in first iteration.
            int u = minDistance(n, dist, sptSet);

            // Mark the picked vertex as processed
            sptSet[u] = true;

            // Update dist value of the adjacent
            // vertices of the picked vertex.
            if (this.Children.TryGetValue(u, out List<(int, int)> children))
            {
                foreach(var childDist in children)
                {
                    var v = childDist.Item1;
                    var uvDist = childDist.Item2;
                    // Update dist[v] only if is not in
                    // sptSet, there is an edge from u
                    // to v, and total weight of path
                    // from src to v through u is smaller
                    // than current value of dist[v]
                    if (!sptSet[v] && dist[u] != int.MaxValue
                        && dist[u] + uvDist < dist[v])
                        dist[v] = dist[u] + uvDist;
                }
            }
        }
        return dist;
    }

    // A utility function to find the
    // vertex with minimum distance
    // value, from the set of vertices
    // not yet included in shortest
    // path tree
    private int minDistance(int n, int[] dist, bool[] sptSet)
    {
        // Initialize min value
        int min = int.MaxValue, min_index = -1;

        for (int v = 0; v < n; v++)
            if (sptSet[v] == false && dist[v] <= min) {
                min = dist[v];
                min_index = v;
            }

        return min_index;
    }

    public int? MinDistanceBFS(int from, int to)
    {
        if (from == to)
        {
            return 0;
        }
        var queue = new Queue<int>();
        var visited = new Dictionary<int, bool>();
        var dist = new Dictionary<int, int>();

        queue.Enqueue(from);
        visited[from] = true;
        dist[from] = 0;

        while(queue.Count > 0)
        {
            var popped = queue.Dequeue();

            if (this.Children.TryGetValue(popped, out List<(int, int)> children))
            {
                foreach (var childDist in children)
                {
                    var child = childDist.Item1;
                    if (child == to)
                    {
                        return dist[popped] + 1;
                    }
                    if (!visited.GetValueOrDefault(child))
                    {
                        queue.Enqueue(child);
                        visited[child] = true;
                        dist[child] = dist[popped] + 1;
                    }
                }
            }
        }
        return null;
    }

    public List<int> EulerianPath(int curr_v = -1)
    {
        if (this.Children.Count == 0)
        {
            return new List<int>();
        }
        var newChildren = this.Children.ToDictionary(entry => entry.Key, entry => entry.Value);
        // edgeCount represents the number of edges emerging from a vertex
        Dictionary<int, int> edgeCount = new Dictionary<int, int>();
        foreach(var link in newChildren) 
        {    
            // find the count of edges to keep track
            // of unused edges
            edgeCount[link.Key] = link.Value.Count;

            if (curr_v == -1)
            {
                curr_v = link.Key; // any vertex will work
            }
        }
        // Maintain a stack to keep vertices
        Stack<int> currPath = new Stack<int>();
    
        // vector to store final circuit
        List<int> circuit = new List<int>();

        // start from any vertex
        currPath.Push(curr_v);
        while (currPath.Count != 0) 
        {    
            // If there's remaining edge
            if (edgeCount.ContainsKey(curr_v) && edgeCount[curr_v] != 0)
            {        
                // Push the vertex
                currPath.Push(curr_v);

                int next_v = this.Children[curr_v].Last().Item1;
                this.Children[curr_v].RemoveAt(this.Children[curr_v].Count - 1);
        
                // and remove that edge
                edgeCount[curr_v]--;
        
                // Move to next vertex
                curr_v = next_v;
            }
        
            // back-track to find remaining circuit
            else
            {
                circuit.Add(curr_v);
        
                // Back-tracking
                curr_v = currPath.Pop();
            }
        }

        var result = new List<int>();
    
        // we've got the circuit, now print it in reverse
        for (int i = circuit.Count - 1; i >= 0; i--) 
        {
            result.Add(circuit[i]);
        }
        return result;
    }
}

// https://leetcode.com/problems/design-graph-with-shortest-path-calculator/
type To struct {
    Node int
    Cost int
}

type Graph struct {
    Links map[int][]To
    N int
}


func Constructor(n int, edges [][]int) Graph {
    result := Graph{}
    result.Links = make(map[int][]To)
    for _, link := range edges {
        result.Links[link[0]] = append(result.Links[link[0]], To{Node: link[1], Cost: link[2]})
    }
    result.N = n
    return result
}


func (this *Graph) AddEdge(edge []int)  {
    this.Links[edge[0]] = append(this.Links[edge[0]], To{Node: edge[1], Cost: edge[2]})
}


func (this *Graph) ShortestPath(node1 int, node2 int) int {
    // dijkstra algo    
    n := this.N
    dist := make([]int, n)
    sptSet := make([]bool, n)

    for i := 0; i < n; i++ {
        dist[i] = math.MaxInt
    }

    dist[node1] = 0

    for count := 0; count < n - 1; count++ {
            // Pick the minimum distance vertex
            // from the set of vertices not yet
            // processed. u is always equal to
            // src in first iteration.
            u := minDistance(dist, sptSet, n)
 
            // Mark the picked vertex as processed
            sptSet[u] = true
 
            // Update dist value of the adjacent
            // vertices of the picked vertex.

            for _, v := range this.Links[u] { 
                // Update dist[v] only if is not in
                // sptSet, there is an edge from u
                // to v, and total weight of path
                // from src to v through u is smaller
                // than current value of dist[v]
                if !sptSet[v.Node] && dist[u] != math.MaxInt && dist[u] + v.Cost < dist[v.Node] {
                    dist[v.Node] = dist[u] + v.Cost
                }                
            }
        }
    if dist[node2] == math.MaxInt {
        return -1
    }
    return dist[node2]
}

func minDistance(dist []int, sptSet []bool, n int) int {
    // Initialize min value
    min, min_index := math.MaxInt,  -1
 
    for v := 0; v < n; v++ {
        if !sptSet[v] && dist[v] <= min {
            min = dist[v]
            min_index = v
        }
    }
    return min_index
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}


/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */

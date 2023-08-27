// https://leetcode.com/problems/minimum-genetic-mutation/
func minMutation(startGene string, endGene string, bank []string) int {
    // add startGene to the list of banks, since it can be not there
    updatedBank :=[]string{startGene}
    startGeneIndex := 0
    endGeneIndex := -1
    startExist := false
    for _, gene := range bank {
        if gene == startGene {
            continue
        }
        if gene == endGene {
            endGeneIndex = len(updatedBank)
        }
        updatedBank = append(updatedBank, gene)
    }
    if !startExist {
        updatedBank = append(updatedBank, startGene)
    }

    bankMap := map[string]bool{}
    links := make([][]int, len(updatedBank))
    for _, gene := range updatedBank {
        bankMap[gene] = true
    }

    // create graph of genes
    for i, gene := range updatedBank {
        for j := i+1; j < len(updatedBank); j++ {
            if areNeighbours(gene, updatedBank[j]) {
                links[i] = append(links[i], j)
                links[j] = append(links[j], i)
            }
        }
    }
    // use regular BFS approach to iternate from start to end gene
    visited := map[int]bool{}
    visited[startGeneIndex] = true

    jumpCount := map[int]int{}
    jumpCount[startGeneIndex] = 0

    queue := []int{startGeneIndex}

    for len(queue) > 0 {
        popped := queue[0]
        queue = queue[1:]
        poppedJumpCount := jumpCount[popped]

        for _, child := range links[popped] {
            if child == endGeneIndex {
                return poppedJumpCount + 1
            }
            if !visited[child] {
                queue = append(queue, child)
                visited[child] = true
                jumpCount[child] = poppedJumpCount + 1
            }
        }
    }
    return -1
}

// returns whether two genes are neighbours
func areNeighbours(s1, s2 string) bool {
    diffFound := false
    for i := range s1 {
        if s1[i] != s2[i] {
            if diffFound {
                return false
            }
            diffFound = true
        }
    }
    return diffFound
}

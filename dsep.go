package dsep

import "errors"

type node struct {
	parents []int
	childs  []int
}

type dVertex struct {
	id   int
	isUp bool
}

// FindDSeperation finds all D-seperation relationship from source node given observations
func FindDSeperation(adjList [][]int, src int, obs []int) ([]int, error) {
	var seperated []int
	if isInSlice(src, obs) {
		return seperated, errors.New("source node should not be in the observation nodes")
	}

	obsDict := make(map[int]bool)
	for _, vertex := range obs {
		obsDict[vertex] = true
	}

	reachable := findReachable(adjList, src, obsDict)

	for i := 0; i < len(adjList); i++ {
		if !reachable[i] && !obsDict[i] && i != src {
			seperated = append(seperated, i)
		}
	}
	return seperated, nil
}

func findReachable(adjList [][]int, src int, obs map[int]bool) map[int]bool {
	// convert the adjList to customized graph
	bys := toBysNet(adjList)

	// ====== Phase 1, Find all ancestors of observations ======
	ancestors := make(map[int]bool)

	// copy obs to queue
	queue := make([]int, len(obs)) // nodes to be visited
	for k := range obs {
		queue = append(queue, k)
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:len(queue)]
		if !ancestors[node] {
			queue = append(queue, bys[node].parents...)
			ancestors[node] = true
		}
	}

	// ====== Phase 2, traverse active trails starting from source ======
	visited := make(map[dVertex]bool)       // vististed nodes
	reachable := make(map[int]bool)         // reachable nodes
	dQueue := []dVertex{dVertex{src, true}} // queue with direction

	for len(dQueue) > 0 {
		dNode := dQueue[0]
		dQueue = dQueue[1:len(dQueue)]
		if !visited[dNode] {
			if !obs[dNode.id] {
				reachable[dNode.id] = true
			}
			visited[dNode] = true
			if dNode.isUp && !obs[dNode.id] {
				for _, node := range bys[dNode.id].parents {
					dQueue = append(dQueue, dVertex{node, true})
				}
				for _, node := range bys[dNode.id].childs {
					dQueue = append(dQueue, dVertex{node, false})
				}
			} else if !dNode.isUp {
				if !obs[dNode.id] {
					for _, node := range bys[dNode.id].childs {
						dQueue = append(dQueue, dVertex{node, false})
					}
				}
				if ancestors[dNode.id] {
					for _, node := range bys[dNode.id].parents {
						dQueue = append(dQueue, dVertex{node, true})
					}
				}
			}
		}
	}
	return reachable
}

func toBysNet(adjList [][]int) []node {
	bys := make([]node, len(adjList))
	for i, row := range adjList {
		for _, v := range row {
			bys[i].childs = append(bys[i].childs, v)
			bys[v].parents = append(bys[v].parents, i)
		}
	}
	return bys
}

func isInSlice(target int, s []int) bool {
	for _, v := range s {
		if target == v {
			return true
		}
	}
	return false
}

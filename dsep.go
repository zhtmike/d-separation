package dsep

import "errors"

type node struct {
	parents []int
	childs  []int
}

type bysNet []node

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
	reachable := findReachable(adjList, src, obs)

	for i := 0; i < len(adjList); i++ {
		if !isInSlice(i, reachable) && !isInSlice(i, obs) && i != src {
			seperated = append(seperated, i)
		}
	}
	return seperated, nil
}

func findReachable(adjList [][]int, src int, obs []int) []int {
	// convert the adjList to customized graph
	bys := toBysNet(adjList)

	// ====== Phase 1, Find all ancestors of observations ======
	var ancestors []int
	queue := make([]int, len(obs)) // nodes to be visited
	copy(queue, obs)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:len(queue)]
		if !isInSlice(node, ancestors) {
			queue = append(queue, bys[node].parents...)
			ancestors = append(ancestors, node)
		}
	}

	// ====== Phase 2, traverse active trails starting from source ======
	var visited []dVertex                   // vististed nodes
	var reachable []int                     // reachable nodes
	dQueue := []dVertex{dVertex{src, true}} // queue with direction

	for len(dQueue) > 0 {
		dNode := dQueue[0]
		dQueue = dQueue[1:len(dQueue)]
		if !isInDSlice(dNode, visited) {
			if !isInSlice(dNode.id, obs) && !isInSlice(dNode.id, reachable) {
				reachable = append(reachable, dNode.id)
			}
			visited = append(visited, dNode)
			if dNode.isUp && !isInSlice(dNode.id, obs) {
				for _, node := range bys[dNode.id].parents {
					dQueue = append(dQueue, dVertex{node, true})
				}
				for _, node := range bys[dNode.id].childs {
					dQueue = append(dQueue, dVertex{node, false})
				}
			} else if !dNode.isUp {
				if !isInSlice(dNode.id, obs) {
					for _, node := range bys[dNode.id].childs {
						dQueue = append(dQueue, dVertex{node, false})
					}
				}
				if isInSlice(dNode.id, ancestors) {
					for _, node := range bys[dNode.id].parents {
						dQueue = append(dQueue, dVertex{node, true})
					}
				}
			}
		}
	}
	return reachable
}

func toBysNet(adjList [][]int) bysNet {
	bys := make(bysNet, len(adjList))
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

func isInDSlice(target dVertex, s []dVertex) bool {
	for _, v := range s {
		if target == v {
			return true
		}
	}
	return false
}

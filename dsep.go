// Package dsep finds all D-separeted nodes from a given node in O(n) time complexity.
//
// Reference: Daphne Koller and Nir Friedman, Probabilistic Graphical Models Principles and Techniques, p74-75.
package dsep

import "errors"

// node stores the parents and childs for a given vertex
type node struct {
	parents []int
	childs  []int
}

// dVertex stores the vertex's id and its traversal direction in the second phase in node searching
type dVertex struct {
	id   int
	isUp bool
}

// FindDSeperation returns all D-seperation nodes from a source node given a observation set.
// A node X is D-separeted from Y given obsrvation Z implies X is conditionally independent to Y given Z.
//
// AdjList must be labelled from 0 to N-1 in adjacent list format, i.e
// [[1], [], []] means there is directed edge from 0 to 1, and there is no outgoing edge from 1 or 2.
func FindDSeperation(adjList [][]int, src int, obs []int) ([]int, error) {
	var seperated []int

	if isInSlice(src, obs) {
		return seperated, errors.New("source node should not be in the observation nodes")
	} else if src > len(adjList) || src < 0 {
		return seperated, errors.New("source node is not in the adjlist")
	}

	// convert observation from slice to dict
	obsDict := make(map[int]bool)
	for _, vertex := range obs {
		obsDict[vertex] = true
	}
	// convert the adjList to node slice
	bys := toBysNet(adjList)

	// find the reachable nodes from src
	reachable := findReachable(bys, src, obsDict)

	// convert the d-separated nodes from dict to slice
	for i := 0; i < len(adjList); i++ {
		if !reachable[i] && !obsDict[i] && i != src {
			seperated = append(seperated, i)
		}
	}
	return seperated, nil
}

func findReachable(bys []node, src int, obs map[int]bool) map[int]bool {
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

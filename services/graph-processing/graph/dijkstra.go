package graph

import (
	"container/heap"
	"math"
)

type Item struct {
	Node     string
	Distance float64
	Path     []string
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].Distance < pq[j].Distance }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i]; pq[i].Index = i; pq[j].Index = j }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Item)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Dijkstra(g *Graph, source, target string) ([]string, float64) {
	dist := map[string]float64{}
	prev := map[string]string{}
	for node := range g.Nodes {
		dist[node] = math.Inf(1)
	}
	dist[source] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{Node: source, Distance: 0, Path: []string{source}})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Item)
		if curr.Node == target {
			return curr.Path, curr.Distance
		}

		for _, edge := range g.Nodes[curr.Node] {
			alt := curr.Distance + edge.Distance
			if alt < dist[edge.Target] {
				dist[edge.Target] = alt
				prev[edge.Target] = curr.Node
				heap.Push(pq, &Item{
					Node:     edge.Target,
					Distance: alt,
					Path:     append(append([]string{}, curr.Path...), edge.Target),
				})
			}
		}
	}
	return nil, math.Inf(1)
}

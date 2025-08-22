package graph

import	"github.com/oceakun/chariotx/services/graph-processing/models"

type Graph struct {
	Nodes map[string][]Edge
}

type Edge struct {
	Target   string
	Distance float64
}

func BuildGraph(segments []models.Segment) *Graph {
	graph := &Graph{Nodes: make(map[string][]Edge)}

	for _, seg := range segments {
		graph.Nodes[seg.Source] = append(graph.Nodes[seg.Source], Edge{
			Target:   seg.Target,
			Distance: seg.Distance,
		})
	}
	return graph
}

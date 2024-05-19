package graph

import (
	"fmt"
	"strings"
)

type AdjacencyList struct {
  Vertices    int
  Edges       int
  List        [][]Edge
}

func NewAdjacencyList(vertices int, edges []Edge) (*AdjacencyList, error) {
  if vertices <= 0 {
    return nil, fmt.Errorf("Number of vertices must be positive")
  }

  if len(edges) < 1 {
    return nil, fmt.Errorf("Number of edges must be at least equal 1")
  }

  edgesLength := len(edges)

  al := &AdjacencyList{Vertices: vertices, Edges: edgesLength}
  al.List = make([][]Edge, vertices)

  for i := 0; i < edgesLength; i++ {
    if edges[i].Source < 0 || edges[i].Source > al.Vertices {
      return nil, fmt.Errorf("Invalid source vertex in edge %v", edges[i])
    }

    if edges[i].Destination < 0 || edges[i].Destination > al.Vertices {
      return nil, fmt.Errorf("Invalid destination vertex in edge %v", edges[i])
    }

    al.List[edges[i].Source] = append(al.List[edges[i].Source], edges[i])
  }

  return al, nil
}

func (al *AdjacencyList) GetVerticesNumber() int {
  return al.Vertices
}

func (al *AdjacencyList) GetEdgesNumber() int {
  return al.Edges
}

func (al *AdjacencyList) LookupEdge(source, destination int) (bool, Edge, error) {
  if source < 0 || destination < 0 || source >= al.Vertices || destination >= al.Vertices {
    return false, Edge{}, fmt.Errorf("Invalid source or destination (%v, %v)", source, destination)
  }

  for i := 0; i < len(al.List[source]); i++ {
    if al.List[source][i].Destination == destination {
      return true, al.List[source][i], nil
    }
  }

  return false, Edge{}, nil
}

func (al *AdjacencyList) GetNeighbors(vertex int) ([]Edge, error) {
  if vertex < 0 || vertex >= al.Vertices {
    return nil, fmt.Errorf("Invalid vertex: %v", vertex)
  }

  neighbors := make([]Edge, len(al.List[vertex]))
  copy(neighbors, al.List[vertex])

  return neighbors, nil
}

func (al *AdjacencyList) GetEdges() []Edge {
  edges := make([]Edge, al.Edges)

  edgeIdx := 0
  for v := 0; v < al.Vertices; v++ {
    edgeCount := len(al.List[v])
    for e := 0; e < edgeCount && edgeIdx < al.Edges; e++ {
      edges[edgeIdx] = al.List[v][e]
      edgeIdx++
    }
  }

  return edges
}

func (al *AdjacencyList) ToString() string {
  var output []string
  for i := 0; i < al.Vertices; i++ {
    output = append(output, fmt.Sprintf("%v", al.List[i]))
  }

  return strings.Join(output, "\n")
}

package graph

import (
	"fmt"
	"strings"
)

type IncidenceMatrix struct {
  Vertices  int
  Edges     int
  Weights   []int
  Matrix    [][]int
}

func NewIncidenceMatrix(vertices int, edges []Edge) (*IncidenceMatrix, error) {
  if vertices <= 0 {
    return nil, fmt.Errorf("Number of vertices must be positive")
  }

  if len(edges) < 1 {
    return nil, fmt.Errorf("Number of edges must be at least equal 1")
  }

  im := &IncidenceMatrix{Vertices: vertices, Edges: len(edges)}
  im.Matrix = make([][]int, vertices)

  edgesLength := len(edges)
  for i := 0; i < vertices; i++ {
    im.Matrix[i] = make([]int, edgesLength)
  }

  im.Weights = make([]int, edgesLength)
  for i := 0; i < edgesLength; i++ {
    if edges[i].Source < 0 || edges[i].Source >= im.Vertices {
      return nil, fmt.Errorf("Invalid source vertex in edge %v", edges[i])
    }

    if edges[i].Destination < 0 || edges[i].Destination >= im.Vertices {
      return nil, fmt.Errorf("Invalid destination vertex in edge %v", edges[i])
    }
    
    im.Matrix[edges[i].Source][i] = 1
    im.Matrix[edges[i].Destination][i] = -1
    im.Weights[i] = edges[i].Weight
  }

  return im, nil

}

func (im *IncidenceMatrix) GetVerticesNumber() int {
  return im.Vertices
}

func (im *IncidenceMatrix) GetEdgesNumber() int {
  return im.Edges
}

func (im *IncidenceMatrix) LookupEdge(source, destination int) (bool, Edge, error) {
  if source < 0 || destination < 0 || source >= im.Vertices || destination >= im.Vertices {
    return false, Edge{}, fmt.Errorf("Invalid source or destination (%v, %v)", source, destination)
  }

  for i := 0; i < im.Edges; i++ {
    if im.Matrix[source][i] == 1 && im.Matrix[destination][i] == -1 {
      return true, Edge{Source: source, Destination: destination, Weight: im.Weights[i]}, nil
    }
  }

  return false, Edge{}, nil
}

func (im *IncidenceMatrix) GetNeighbors(vertex int) ([]Edge, error) {
  if vertex < 0 || vertex >= im.Vertices {
    return nil, fmt.Errorf("Invalid vertex: %v", vertex)
  }

  var neighbors []Edge
  for i := 0; i < im.Vertices; i++ {
    if vertex == i {
      continue
    }
    if exists, edge, err := im.LookupEdge(vertex, i); err == nil && exists {
      neighbors = append(neighbors, edge)
    }
  }

  return neighbors, nil
}

func (im *IncidenceMatrix) ToString() string {
  var output []string
  for i := 0; i < im.Vertices; i++ {
    output = append(output, fmt.Sprintf("%v", im.Matrix[i]))
  }

  output = append(output, fmt.Sprintf("%v", im.Weights))
  return strings.Join(output, "\n")
}

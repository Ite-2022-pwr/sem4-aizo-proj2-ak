package graph

import (
	"fmt"
	"math"
)

func minDistance(distances []int, visited []bool) int {
  verticesNumber := len(visited)
  var minDistanceVertex int
  minDist := math.MaxInt

  for v := 0; v < verticesNumber; v++ {
    if !visited[v] && distances[v] < minDist {
      minDist, minDistanceVertex = distances[v], v
    }
  }

  return minDistanceVertex
}

func Dijkstra(G Graph, startVertex int) (distances []int, parents []int, err error) {
  if startVertex < 0 || startVertex >= G.GetVerticesNumber() {
    return nil, nil, fmt.Errorf("Invalid start vertex number: %v", startVertex)
  }

  distances = make([]int, G.GetVerticesNumber())
  visited := make([]bool, G.GetVerticesNumber())
  parents = make([]int, G.GetVerticesNumber())

  for i := 0; i < G.GetVerticesNumber(); i++ {
    distances[i] = math.MaxInt
  }

  distances[startVertex] = 0

  for i := 0; i < G.GetVerticesNumber() - 1; i++ {
    u := minDistance(distances, visited)

    visited[u] = true

    for v := 0; v < G.GetVerticesNumber(); v++ {
      if visited[v] {
        continue
      }

      exists, edge, err := G.LookupEdge(u, v)
      if err != nil {
        return nil, nil, err
      }

      if !exists {
        continue
      }

      if distances[u] + edge.Weight < distances[v] {
        distances[v] = distances[u] + edge.Weight
        parents[v] = u
      }
    }
  }

  return distances, parents, nil
}


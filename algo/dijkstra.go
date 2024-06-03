package algo

import (
	"fmt"
	"math"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/graph"
)

// FindPathString znajduje dokładną ścieżkę od jednego wierzchołka do drugiego
// na podstawie tablicy parents
func FindPathString(parents []int, begin, end int) (string, error) {
  if begin < 0 || begin > len(parents) - 1 {
    return "", fmt.Errorf("Invalid begin: %v", begin)
  }
  if end < 0 || end > len(parents) - 1 {
    return "", fmt.Errorf("Invalid end: %v", end)
  }

  ret := fmt.Sprintf("-> %v", end)
  v := end
  for v != begin {
    v = parents[v]
    ret = fmt.Sprintf("-> %v ", v) + ret
  }

  return ret, nil
}

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

// Dijkstra znajduje najkrótsze ścieżki z podanego wierzchołka do wszystkich pozostałych
// za pomocą naiwnej implementacji algorytmu Dijkstry
func Dijkstra(G graph.Graph, startVertex int) (distances []int, parents []int, err error) {
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

      exists, edge, err := G.LookupEdge(u, v, true)
      if err != nil {
        return nil, nil, err
      }

      if !exists {
        continue
      }

      // relaksacja krawędzi
      if distances[u] + edge.Weight < distances[v] {
        distances[v] = distances[u] + edge.Weight
        parents[v] = u
      }
    }
  }

  return distances, parents, nil
}


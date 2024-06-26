package algo

import (
	"fmt"
	"math"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/graph"
)

// BellmanFord znajduje najkrótsze ścieżki od podanego wierchołka do wszystkich pozostałych wierchołków
// w grafie za pomocą algorytmu Forda-Bellmana
func BellmanFord(G graph.Graph, startVertex int) (distances []int, parents []int, err error) {
  if startVertex < 0 || startVertex >= G.GetVerticesNumber() {
    return nil, nil, fmt.Errorf("Invalid start vertex number: %v", startVertex)
  }

  distances = make([]int, G.GetVerticesNumber())
  parents = make([]int, G.GetVerticesNumber())

  for i := 0; i < G.GetVerticesNumber(); i++ {
    distances[i] = math.MaxInt
  }

  distances[startVertex] = 0

  edges := G.GetEdges()

  for i := 0; i < G.GetVerticesNumber() - 1; i++ {
    hasChanged := false
    for e := 0; e < G.GetEdgesNumber(); e++ {
      edge := edges[e]
      u, v := edge.Source, edge.Destination

      // relaksacja krawędzi
      if distances[u] + edge.Weight < distances[v] {
        distances[v] = distances[u] + edge.Weight
        parents[v] = u
        hasChanged = true
      }
    }

    if !hasChanged {
      break
    }
  }

  for e := 0; e < G.GetEdgesNumber(); e++ {
    edge := edges[e]
    u, v := edge.Source, edge.Destination

     if distances[u] + edge.Weight < distances[v] {
       return nil, nil, fmt.Errorf("Graph contains negative cycle")
     }
   }

  return distances, parents, nil
}

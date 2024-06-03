package algo

import (
	"math"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/graph"
)

func findMinKey(keys []int, mstSet []bool) int {
  V, minKey := len(keys), math.MaxInt

  var minKeyIdx int
  for v := 0; v < V; v++ {
    if !mstSet[v] && keys[v] < minKey {
      minKey, minKeyIdx = keys[v], v
    }
  }

  return minKeyIdx
}

// Prim znajduje minimalne drzewo rozpinające w grafie algorytmem Prima
func Prim(G graph.Graph) (mstWeight int, mstEdges []graph.Edge, err error) {

  parents := make([]int, G.GetVerticesNumber())
  mstSet := make([]bool, G.GetVerticesNumber())
  keys := make([]int, G.GetVerticesNumber())

  for i := 0; i < G.GetVerticesNumber(); i++ {
    keys[i] = math.MaxInt
  }

  keys[0] = 0

  for i := 0; i < G.GetVerticesNumber(); i++ {
    u := findMinKey(keys, mstSet)

    mstSet[u] = true

    for v := 0; v < G.GetVerticesNumber(); v++ {
      if mstSet[v] {
        continue
      }

      exists, edge, err := G.LookupEdge(u, v, false)

      if err != nil {
        return 0, nil, err
      }

      if !exists {
        continue
      }

      if edge.Weight < keys[v] {
        parents[v] = u
        keys[v] = edge.Weight
      }
    }
  }

  mstWeight = 0
  for _, v := range keys {
    mstWeight += v
  }

  for v, u := range parents {
    if v == u {
      continue
    }
    mstEdges = append(mstEdges, graph.Edge{Source: u, Destination: v, Weight: keys[v]})
  }

  return mstWeight, mstEdges, nil
}

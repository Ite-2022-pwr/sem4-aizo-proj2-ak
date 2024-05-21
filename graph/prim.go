package graph

import (
	"fmt"
	"math"
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

func Prim(G Graph, startVertex int) (mstWeight int, parents []int, err error) {
  if startVertex < 0 || startVertex >= G.GetVerticesNumber() {
    return 0, nil, fmt.Errorf("Invalid start vertex")
  }

  mstWeight = 0
  parents = make([]int, G.GetVerticesNumber())
  mstSet := make([]bool, G.GetVerticesNumber())
  keys := make([]int, G.GetVerticesNumber())

  for i := 0; i < G.GetVerticesNumber(); i++ {
    keys[i] = math.MaxInt
  }

  keys[startVertex] = 0

  for i := 0; i < G.GetVerticesNumber(); i++ {
    u := findMinKey(keys, mstSet)

    mstSet[u] = true

    for v := 0; v < G.GetVerticesNumber(); v++ {
      if mstSet[v] {
        continue
      }

      exists, edge, err := G.LookupEdge(u, v)

      if err != nil {
        return 0, nil, err
      }

      if !exists {
        continue
      }

      if edge.Weight < keys[v] {
        parents[v] = u
        if keys[v] != math.MaxInt {
          mstWeight -= keys[v]
        }
        keys[v] = edge.Weight
        mstWeight += edge.Weight
      }
    }
  }

  return mstWeight, parents, nil
}

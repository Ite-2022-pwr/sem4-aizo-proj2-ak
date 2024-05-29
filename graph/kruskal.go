package graph

import (
	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/utils"
)

type UnionFind struct {
  Parents []int
  Sizes   []int
}

func NewUnionFind(numberOfVertices int) *UnionFind {
  parents := make([]int, numberOfVertices)
  sizes := make([]int, numberOfVertices)

  for i := 0; i < numberOfVertices; i++ {
    parents[i] = i
    sizes[i] = 1
  }

  return &UnionFind{Parents: parents, Sizes: sizes}
}

func (uf *UnionFind) Find(v int) int {
  for v != uf.Parents[v] {
    uf.Parents[v] = uf.Parents[uf.Parents[v]]
    v = uf.Parents[v]
  }
  return v
}

func (uf *UnionFind) Union(u, v int) bool {
  rootU, rootV := uf.Find(u), uf.Find(v)

  if rootU == rootV {
    return false
  }

  if uf.Sizes[rootU] > uf.Sizes[rootV] {
    uf.Parents[rootV] = rootU
    uf.Sizes[rootU]++
  } else {
    uf.Parents[rootU] = rootV
    uf.Sizes[rootV]++
  }

  return true
}

func CompareEdges(e1, e2 Edge) int {
  if e1.Weight > e2.Weight {
    return 1
  } else if e1.Weight < e2.Weight {
    return -1
  }
  return 0
}

func Kruskal(G Graph) (mstWeight int, mstEdges []Edge) {
  mstWeight = 0

  edges := G.GetEdges()
  uf := NewUnionFind(G.GetVerticesNumber())
  edgeCount := 0

  utils.HeapSort(edges, CompareEdges)

  for _, edge := range edges {
    if uf.Union(edge.Source, edge.Destination) {
      mstWeight += edge.Weight
      edgeCount++
      mstEdges = append(mstEdges, edge)
      if edgeCount == G.GetVerticesNumber() - 1 {
        return mstWeight, mstEdges
      }
    }
  }

  return mstWeight, mstEdges
} 

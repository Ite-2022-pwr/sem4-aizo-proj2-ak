package algo

import (
	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/graph"
	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/utils"
)

// UnionFind reprezentuje rozłączne potrzebne do działania algorytmu Kruskala
type UnionFind struct {
  Parents []int
  Sizes   []int
}

// NewUnionFind zwraca nowy UnionFind
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

// Union łączy zbiory dla dwóch różnych wierzchołków
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

// Kruskal znajduje minimalne drzewo rozpinające w grafie algorytmem Kruskala
func Kruskal(G graph.Graph) (mstWeight int, mstEdges []graph.Edge) {
  mstWeight = 0

  edges := G.GetEdges()
  uf := NewUnionFind(G.GetVerticesNumber())
  edgeCount := 0

  utils.HeapSort(edges, graph.CompareEdges)

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

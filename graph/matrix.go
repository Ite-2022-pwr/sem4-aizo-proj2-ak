package graph

import (
	"fmt"
	"strings"
)

// AdjacencyMatrix to reprezentacja grafu w formie
// macierzy incydencji.
type AdjacencyMatrix struct {
  VerticesNumber int
  Matrix [][]int
}

func NewAdjacencyMatrix(vertices int) *AdjacencyMatrix {
  am := AdjacencyMatrix{}
  am.AddVertices(vertices)
  return &am
}

// AddVertices dodaje n nowych wierzchołków do grafu.
func (am *AdjacencyMatrix) AddVertices(n int) error {
  if n <= 0 {
    return fmt.Errorf("Ujemna liczba nowych wierzchołków")
  }

  matrix := make([][]int, am.VerticesNumber + n)
  for i := 0; i < am.VerticesNumber + n; i++ {
    matrix[i] = make([]int, am.VerticesNumber + n)
  }

  for i := 0; i < am.VerticesNumber; i++ {
    for j := 0; j < am.VerticesNumber; j++ {
      matrix[i][j] = am.Matrix[i][j]
    }
  }

  am.VerticesNumber += n
  am.Matrix = matrix

  return nil
}

func (am *AdjacencyMatrix) AddEdge(u, v, cost int) error {
  if u < 0 || v < 0 {
    return fmt.Errorf("Numery wierzchołków muszą być dodatnie")
  }

  if u > am.VerticesNumber || v > am.VerticesNumber {
    return fmt.Errorf("Złe numery wierzchołków")
  }

  if u == v {
    return fmt.Errorf("Numery wierzchołków takie same")
  }

  am.Matrix[u][v] = cost
  am.Matrix[v][u] = cost

  return nil
}

func (am *AdjacencyMatrix) ToString() string {
  if am.Matrix == nil {
    return ""
  }

  str := ""

  // str += "  "
  // for i := 0; i < am.VerticesNumber; i++ {
  //   str += fmt.Sprintf("%v ", i)
  // }
  // str = strings.TrimRight(str, " ")
  // str += "\n"

  for i := 0; i < am.VerticesNumber; i++ {
    // str += fmt.Sprintf("%v ", i)
    for j := 0; j <len(am.Matrix[i]); j++ {
      str += fmt.Sprintf("%v ", am.Matrix[i][j])
    }
    str = strings.TrimRight(str, " ")
    str += "\n"
  }

  return strings.TrimRight(str, "\n")
}

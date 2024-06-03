package analysis

import (
	"fmt"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/graph"
)

type GraphAnalyzer struct {
  List *graph.AdjacencyList
  Matrix *graph.IncidenceMatrix
}

func NewGraphAnalyzer(numberOfVertices int, edges []graph.Edge) (ga *GraphAnalyzer, err error) {
  if numberOfVertices < 1 {
    return nil, fmt.Errorf("Number of vertices must be positive, given %v", numberOfVertices)
  }

  al, err := graph.NewAdjacencyList(numberOfVertices, edges)

  if err != nil {
    return nil, err
  }

  im, err := graph.NewIncidenceMatrix(numberOfVertices, edges)

  if err != nil {
    return nil, err
  }

  return &GraphAnalyzer{List: al, Matrix: im}, nil
}

func (ga *GraphAnalyzer) PrintAdjacencyList() {
  fmt.Println("Lista sąsiedztwa:")
  fmt.Println(ga.List.ToString())
}

func (ga *GraphAnalyzer) PrintIncidenceMatrix() {
  fmt.Println("Macierz incydencji:")
  fmt.Println(ga.Matrix.ToString())
}

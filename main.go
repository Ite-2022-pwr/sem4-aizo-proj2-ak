package main

import (
	"fmt"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/graph"
)

func main() {
  G := graph.NewAdjacencyMatrix(6)
  G.AddEdge(0, 1, 1)
  G.AddEdge(0, 4, 1)
  G.AddEdge(1, 4, 1)
  G.AddEdge(1, 2, 1)
  G.AddEdge(3, 4, 1)
  G.AddEdge(2, 3, 1)
  G.AddEdge(3, 5, 1)

  fmt.Println(G.ToString())
}

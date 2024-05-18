package main

import (
	"fmt"
	"log"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/graph"
)

func main() {
  edges := []graph.Edge{
    {Source: 0, Destination: 1, Weight: 3},
    {Source: 0, Destination: 4, Weight: 3},
    {Source: 1, Destination: 2, Weight: 1},
    {Source: 2, Destination: 3, Weight: 3},
    {Source: 2, Destination: 5, Weight: 1},
    {Source: 3, Destination: 1, Weight: 3},
    {Source: 4, Destination: 5, Weight: 2},
    {Source: 5, Destination: 0, Weight: 6},
    {Source: 5, Destination: 3, Weight: 1},
  }

  var G graph.Graph
  var err error

  G, err = graph.NewIncidenceMatrix(6, edges)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(G.ToString())
  fmt.Println()

  fmt.Println(G.LookupEdge(0, 4))
  fmt.Println(G.LookupEdge(0, 5))

  fmt.Println(G.GetNeighbors(2))

  fmt.Println()

  G, err = graph.NewAdjacencyList(6, edges)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(G.ToString())
  fmt.Println()

  fmt.Println(G.LookupEdge(0, 4))
  fmt.Println(G.LookupEdge(0, 5))

  fmt.Println(G.GetNeighbors(2))
}

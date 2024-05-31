package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/analysis"
	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/generator"
	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/graph"
	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/utils"
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

  // var G graph.Graph
  var err error
  //
  // G, err = graph.NewIncidenceMatrix(6, edges)
  //
  // if err != nil {
  //   log.Fatal(err)
  // }

  // fmt.Println(G.ToString())
  // fmt.Println()

  // fmt.Println(G.LookupEdge(0, 4))
  // fmt.Println(G.LookupEdge(0, 5))

  // fmt.Println(G.GetNeighbors(2))

  // fmt.Println(G.GetEdges())

  // fmt.Println(algo.Dijkstra(G, 0))
  // fmt.Println(algo.BellmanFord(G, 0))
  // fmt.Println(algo.Prim(G))
  // fmt.Println(algo.Kruskal(G))

  // fmt.Println()
  //
  // G, err = graph.NewAdjacencyList(6, edges)
  //
  // if err != nil {
  //   log.Fatal(err)
  // }

  // fmt.Println(G.ToString())
  // fmt.Println()

  // fmt.Println(G.LookupEdge(0, 4))
  // fmt.Println(G.LookupEdge(0, 5))

  // fmt.Println(G.GetNeighbors(2))

  // fmt.Println(G.GetEdges())

  // fmt.Println(algo.Dijkstra(G, 0))
  // fmt.Println(algo.BellmanFord(G, 0))
  // fmt.Println(algo.Prim(G))
  // fmt.Println(algo.Kruskal(G))
  //
  // fmt.Println()
  // fmt.Println()
  // fmt.Println()
  // fmt.Println()
  // fmt.Println()

  fmt.Println("Generating graphs")

  V := 100
  density := 25

  edges, err = generator.GenerateGraph(V, 10, density, false)

  if err != nil {
    fmt.Println(err)
    return
  }

  // al, err := graph.NewAdjacencyList(V, edges)
  // 
  // if err != nil {
  //   fmt.Println(err)
  //   return
  // }
  //
  // // fmt.Println(al.ToString())
  // fmt.Println()
  //
  // im, err := graph.NewIncidenceMatrix(V, edges)
  //
  // if err != nil {
  //   fmt.Println(err)
  //   return
  // }

  // fmt.Println(im.ToString())
  // fmt.Println()
  //
  // fmt.Println("Dijkstra")
  // fmt.Println(algo.Dijkstra(al, 0))
  // fmt.Println(algo.Dijkstra(im, 0))
  // fmt.Println()
  //
  // fmt.Println("BellmanFord")
  // fmt.Println(algo.BellmanFord(al, 0))
  // fmt.Println(algo.BellmanFord(im, 0))
  // fmt.Println()
  //
  // fmt.Println("Prim")
  // fmt.Println(algo.Prim(al))
  // fmt.Println(algo.Prim(im))
  // fmt.Println()
  //
  // fmt.Println("Kruskal")
  // fmt.Println(algo.Kruskal(al))
  // fmt.Println(algo.Kruskal(im))
  // fmt.Println()
  //
  // fmt.Println()
  // fmt.Println()
  // fmt.Println()
  // fmt.Println()
  // fmt.Println()
  // fmt.Println()
  
  graphAnal, err := analysis.NewGraphAnalyzer(V, edges)

  if err != nil {
    fmt.Println(err)
    return
  }

  // graphAnal.Dijstkra(0, V - 1)
  // graphAnal.BellmanFord(0, V - 1)

  fmt.Println()

  // graphAnal.Prim()
  // graphAnal.Kruskal()

  utils.SaveGraph("data/input/test.txt", graphAnal.List)
  _, edges, err = utils.ReadGraphFromFile("data/input/test.txt")
  if err != nil {
    log.Fatal(err)
    return
  }

  start := time.Now()
  V = 1000
  density = 99
  edges, err = generator.GenerateGraph(V, 20, density, true)
  utils.PrintTimeElapsed(start, "Generowanie grafu 1000 wierchołków 99%% gęstości")
  if err != nil {
    log.Fatal(err)
  }
  start = time.Now()
  graphAnal,  err = analysis.NewGraphAnalyzer(V, edges)
  utils.PrintTimeElapsed(start, "Generowanie analizatora grafów na podstawie listy krawędzi")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(graphAnal.List.Vertices, graphAnal.List.Edges)
  
  // graphAnal.Dijstkra(0, 999)
}

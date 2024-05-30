package generator

import (
	"fmt"
	"math/rand"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/graph"
)

func GenerateGraph(numberOfVertices, maxWeight, densityPerentage int) (edges []graph.Edge, err error) {
  if numberOfVertices <= 0 {
    return nil, fmt.Errorf("Number of vertices must be positive, given: %v", numberOfVertices)
  }

  if maxWeight <= 0 {
    return nil, fmt.Errorf("Max weight must be positive, given: %v", maxWeight)
  }

  if densityPerentage <= 0 && densityPerentage >= 100{
    return nil, fmt.Errorf("Density percentage must be positive and less than 100, given: %v", densityPerentage)
  }

  numberOfEdges := (numberOfVertices * (numberOfVertices - 1) / 2) * densityPerentage / 100
  // fmt.Println(numberOfEdges)

  if numberOfEdges < numberOfVertices - 1 {
    return nil, fmt.Errorf("Number of edges too small, increase density perentage")
  }

  edgesCount := numberOfVertices - 1

  edges, adjacencyMatrix, err := GenerateMst(numberOfVertices, maxWeight)
  if err != nil {
    return nil, err
  }

  for edgesCount < numberOfEdges {
    src := rand.Intn(numberOfVertices)
    dest := rand.Intn(numberOfVertices)

    if src == dest || adjacencyMatrix[src][dest] != 0 || adjacencyMatrix[dest][src] != 0 {
      continue
    }

    weight := rand.Intn(maxWeight) + 1

    edges = append(edges, graph.Edge{Source: src, Destination: dest, Weight: weight})
    adjacencyMatrix[src][dest], adjacencyMatrix[dest][src] = weight, weight
    edgesCount++
  }

  return edges, nil
}

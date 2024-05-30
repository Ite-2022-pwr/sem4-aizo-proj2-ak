package generator

import (
	"fmt"
	"math/rand"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/graph"
)

func GenerateMst(numberOfVertices, maxWeight int) (mst []graph.Edge, adjacencyMatrix [][]int, err error) {
  if numberOfVertices <= 0 {
    return nil, nil, fmt.Errorf("Number of vertices must be positive, given: %v", numberOfVertices)
  }

  if maxWeight <= 0 {
    return nil, nil, fmt.Errorf("Max weight must be positive, given: %v", maxWeight)
  }

  mst = make([]graph.Edge, numberOfVertices - 1)

  adjacencyMatrix = make([][]int, numberOfVertices)

  for i := 0; i < numberOfVertices; i++ {
    adjacencyMatrix[i] = make([]int, numberOfVertices)
  }

  for v := 1; v < numberOfVertices; v++ {
    src, dest := v - 1, v
    weight := rand.Intn(maxWeight) + 1

    mst[src] = graph.Edge{Source: src, Destination: dest, Weight: weight}
    adjacencyMatrix[src][dest], adjacencyMatrix[dest][src] = weight, weight
  }

  return mst, adjacencyMatrix, nil
}

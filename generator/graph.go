package generator

import (
	"fmt"
	"math/rand"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/graph"
)

func GenerateGraph(numberOfVertices, maxWeight, densityPerentage int, directed bool) (edges []graph.Edge, err error) {
  if numberOfVertices <= 0 {
    return nil, fmt.Errorf("Number of vertices must be positive, given: %v", numberOfVertices)
  }

  if maxWeight <= 0 {
    return nil, fmt.Errorf("Max weight must be positive, given: %v", maxWeight)
  }

  if densityPerentage <= 0 && densityPerentage >= 100{
    return nil, fmt.Errorf("Density percentage must be positive and less than 100, given: %v", densityPerentage)
  }

  numberOfEdges := numberOfVertices * (numberOfVertices - 1)

  if !directed {
    numberOfEdges /= 2
  }

  numberOfEdges = numberOfEdges * densityPerentage / 100
  // fmt.Println(numberOfEdges)

  if numberOfEdges < numberOfVertices - 1 {
    return nil, fmt.Errorf("Number of edges too small, increase density perentage")
  }

  if densityPerentage > 60 && directed {
    var edges []graph.Edge
    for i := 0; i < numberOfVertices; i++ {
      for j := 0; j < numberOfVertices; j++ {
        if i == j {
          continue
        }

        edges = append(edges, graph.Edge{Source: i, Destination: j, Weight: rand.Intn(maxWeight) + 1})
      }
    }

    for i := 0; i < numberOfVertices * (numberOfVertices - 1) - numberOfEdges; i++ {
      idx := rand.Intn(len(edges))
      edges = append(edges[:idx], edges[idx+1:]...)
    }

    return edges, nil
  }

  edgesCount := numberOfVertices - 1

  edges, adjacencyMatrix, err := GenerateMst(numberOfVertices, maxWeight, directed)
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
    adjacencyMatrix[src][dest] = weight
    if !directed {
      adjacencyMatrix[dest][src] = weight
    }
    edgesCount++
  }

  return edges, nil
}

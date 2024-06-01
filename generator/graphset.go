package generator

import (
	"fmt"
	"log"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/analysis"
	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/utils"
)

func GenerateGraphSet(numberOfGraphs, numberOfVertices, maxWeight, densityPerentage int, directed bool) (graphAnalyzers []*analysis.GraphAnalyzer, err error) {
  log.Println(utils.BlueColor(fmt.Sprintf("[*] Generowanie zestawu %v grafów po %v wierzchołków i %v%% gęstości", numberOfGraphs, numberOfVertices, densityPerentage)))

  for i := 0; i < numberOfGraphs; i++ {
    edges, err := GenerateGraph(numberOfVertices, maxWeight, densityPerentage, directed)
    if err != nil {
      log.Println(utils.RedColor(err))
      return nil, err
    }

    ga, err := analysis.NewGraphAnalyzer(numberOfVertices, edges)
    if err != nil {
      log.Println(utils.RedColor(err))
      return nil, err
    }

    graphAnalyzers = append(graphAnalyzers, ga)
  }

  return graphAnalyzers, nil
}

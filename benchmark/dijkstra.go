package benchmark

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/generator"
	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/utils"
)

func BenchmarkDijsktra() {
  log.Println(utils.BlueColor("[+] Rozpoczynanie testowania algorytmu Dijsktry"))

  for _, density := range Densities {
    var outputList, outputMatrix [][]string
    for _, verticiesNum := range VerticiesNumbers {
      listTimeAvg, matrixTimeAvg := 0.0, 0.0
      graphAnals, err := generator.GenerateGraphSet(NumberOfGraphs, verticiesNum, verticiesNum / 10, density, true)
      if err != nil {
        return
      }

      for i, graphAnal := range graphAnals {
        log.Println(utils.YellowColor(fmt.Sprintf("[+] Dijstkra %v wierzchołków %v%% gęstości - pomiar %v/%v", verticiesNum, density, i+1, NumberOfGraphs)))
        listTime, matrixTime := graphAnal.Dijstkra(0, verticiesNum - 1)
        listTimeAvg += listTime
        matrixTimeAvg += matrixTime
      }

      listTimeAvg /= NumberOfGraphs
      matrixTimeAvg /= NumberOfGraphs
      outputList = append(outputList, []string{fmt.Sprintf("%v", verticiesNum), fmt.Sprintf("%.3f", listTimeAvg)})
      outputMatrix = append(outputMatrix, []string{fmt.Sprintf("%v", verticiesNum), fmt.Sprintf("%.3f", matrixTimeAvg)})
    }

    utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("dijkstra_list_density_%v.csv", density)), outputList)
    utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("dijkstra_matrix_density_%v.csv", density)), outputMatrix)
  }
}

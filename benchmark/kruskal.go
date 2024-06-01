package benchmark

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/generator"
	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/utils"
)

func BenchmarkKruskal() {
  log.Println(utils.BlueColor("[+] Rozpoczynanie testowania algorytmu Bellmana-Forda"))

  for _, density := range Densities {
    var outputList, outputMatrix [][]string
    for _, verticiesNum := range VerticiesNumbers {
      listTimeAvg, matrixTimeAvg := 0.0, 0.0
      graphAnals, err := generator.GenerateGraphSet(NumberOfGraphs, verticiesNum, verticiesNum / 10, density, true)
      if err != nil {
        return
      }

      for i, graphAnal := range graphAnals {
        log.Println(utils.YellowColor(fmt.Sprintf("[+] Prim %v wierzchołków %v%% gęstości - pomiar %v/%v", verticiesNum, density, i+1, NumberOfGraphs)))
        listTime, matrixTime := graphAnal.Kruskal(false)
        listTimeAvg += listTime
        matrixTimeAvg += matrixTime
      }

      listTimeAvg /= NumberOfGraphs
      matrixTimeAvg /= NumberOfGraphs
      outputList = append(outputList, []string{fmt.Sprintf("%v", verticiesNum), fmt.Sprintf("%.3f", listTimeAvg)})
      outputMatrix = append(outputMatrix, []string{fmt.Sprintf("%v", verticiesNum), fmt.Sprintf("%.3f", matrixTimeAvg)})
    }

    utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("kruskal_list_density_%v.csv", density)), outputList)
    utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("kruskal_matrix_density_%v.csv", density)), outputMatrix)
  }
}

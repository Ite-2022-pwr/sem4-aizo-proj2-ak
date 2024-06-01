package analysis

import (
	"fmt"
	"log"
	"time"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/algo"
	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/utils"
)

func (ga *GraphAnalyzer) Kruskal(printMstEdges bool) (listTime, matrixTime float64) {
  prompt := utils.BlueColor("Znajdowanie minimalnego drzewa rozpinającego algorytmem Kruskala dla listy sąsiedztwa")
  log.Println("[*] Rozpoczynanie:", prompt)
  start := time.Now()

  mstWeight, mstEdges := algo.Kruskal(ga.List)

  listTime = utils.PrintTimeElapsed(start, prompt)

  log.Println(utils.YellowColor(fmt.Sprintf("[+] Waga minimalnego drzewa rozpinającego: %v", mstWeight)))

  if printMstEdges {
    log.Print(utils.YellowColor(fmt.Sprintf("[+] Krawędzie minimalnego drzewa rozpinajacego")))
    for _, e := range mstEdges {
      fmt.Printf("%v -> %v: %v\n", e.Source, e.Destination, e.Weight)
    }
  }

  prompt = utils.BlueColor("Znajdowanie minimalnego drzewa rozpinającego algorytmem Kruskala dla macierzy incydencji")
  log.Println("[*] Rozpoczynanie:", prompt)
  start = time.Now()

  mstWeight, mstEdges = algo.Kruskal(ga.Matrix)

  matrixTime = utils.PrintTimeElapsed(start, prompt)

  log.Println(utils.YellowColor(fmt.Sprintf("[+] Waga minimalnego drzewa rozpinającego: %v", mstWeight)))

  if printMstEdges {
    log.Print(utils.YellowColor(fmt.Sprintf("[+] Krawędzie minimalnego drzewa rozpinajacego")))
    for _, e := range mstEdges {
      fmt.Printf("%v -> %v: %v\n", e.Source, e.Destination, e.Weight)
    }
  }

  return listTime, matrixTime
}

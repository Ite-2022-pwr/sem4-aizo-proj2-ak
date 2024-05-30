package analysis

import (
	"fmt"
	"log"
	"time"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/algo"
	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/utils"
)

func (ga *GraphAnalyzer) BellmanFord(startVertex, finishVertex int) (listTime, matrixTime float64) {
  if startVertex < 0 || startVertex >= ga.List.Vertices || startVertex >= ga.Matrix.Vertices {
    log.Println(utils.RedColor(fmt.Sprintf("[!!] Nieprawidłowy wierzchołek startowy: %v", startVertex)))
    return 0, 0
  }
  if finishVertex < 0 || finishVertex >= ga.List.Vertices || finishVertex >= ga.Matrix.Vertices {
    log.Println(utils.RedColor(fmt.Sprintf("[!!] Nieprawidłowy wierzchołek końcowy: %v", finishVertex)))
    return 0, 0
  }

  prompt := utils.BlueColor(fmt.Sprintf("Znajdowanie najkrótszej ścieżki algorytmem Bellmana-Forda dla listy sąsiedztwa od wierzchołka %v do %v", startVertex, finishVertex))
  log.Println("[*] Rozpoczynanie:", prompt)
  start := time.Now()

  distances, parents, err := algo.BellmanFord(ga.List, startVertex)

  if err != nil {
    log.Println(utils.RedColor(err))
    return 0, 0
  }

  listTime = utils.PrintTimeElapsed(start, prompt)

  log.Println(utils.YellowColor(fmt.Sprintf("[+] Długość ścieżki od %v do %v: %v", startVertex, finishVertex, distances[finishVertex])))
  path, err := algo.FindPathString(parents, startVertex, finishVertex)
  if err != nil {
    log.Println(utils.RedColor(err))
    return 0, 0
  }
  log.Println(utils.YellowColor(fmt.Sprintf("[+] Ścieżka: %v", path)))

  prompt = utils.BlueColor(fmt.Sprintf("Znajdowanie najkrótszej ścieżki algorytmem Bellmana-Forda dla macierzy incydencji od wierzchołka %v do %v", startVertex, finishVertex))
  log.Println("[*] Rozpoczynanie:", prompt)
  start = time.Now()

  distances, parents, err = algo.BellmanFord(ga.Matrix, startVertex)

  if err != nil {
    log.Println(utils.RedColor(err))
    return 0, 0
  }

  matrixTime = utils.PrintTimeElapsed(start, prompt)

  log.Println(utils.YellowColor(fmt.Sprintf("[+] Długość ścieżki od %v do %v: %v", startVertex, finishVertex, distances[finishVertex])))
  path, err = algo.FindPathString(parents, startVertex, finishVertex)
  if err != nil {
    log.Println(utils.RedColor(err))
    return 0, 0
  }
  log.Println(utils.YellowColor(fmt.Sprintf("[+] Ścieżka: %v", path)))

  return listTime, matrixTime
}

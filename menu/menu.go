package menu

import (
	"fmt"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/analysis"
	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/generator"
	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/utils"
)

const (
  mstProblem = 1
  shortestPathProblem = 2
)

type MenuManager struct {
  Analyzer *analysis.GraphAnalyzer
}

func NewMenuManager() *MenuManager {
  return &MenuManager{}
}

func (mm *MenuManager) Menu() {
  for {
    fmt.Println("Wybierz opcję:")
    fmt.Println("0. Wyjście z programu")
    fmt.Println("1. Wygeneruj graf")
    fmt.Println("2. Wczytaj graf z pliku")
    fmt.Println("3. Wypisz graf jako listę sąsiedztwa")
    fmt.Println("4. Wypisz graf jako macierz incydencji")
    fmt.Println("5. Wykonaj algorytm grafowy")
    fmt.Print("Twój wybór: ")
    var choice int
    if _, err := fmt.Scanln(&choice); err != nil {
      fmt.Println(utils.RedColor(err))
      continue
    }

    switch choice {
    case 0:
      return
    case 1:
      mm.generateGraph()
    case 2:
      mm.readGraphFromFile()
    case 3:
      if mm.Analyzer == nil {
        fmt.Println("Brak grafu do wyświetlenia")
        continue
      }
      mm.Analyzer.PrintAdjacencyList()
    case 4:
      if mm.Analyzer == nil {
        fmt.Println("Brak grafu do wyświetlenia")
        continue
      }
      mm.Analyzer.PrintIncidenceMatrix()
    case 5:
      mm.chooseProblem()
    default:
      fmt.Println("Tylko opcje 0-5")
    }
  }
}

func (mm *MenuManager) generateGraph() {
  var numberOfVertices, density int

  fmt.Print("Podaj liczbę wierzchołków: ")
  if _, err := fmt.Scanln(&numberOfVertices); err != nil {
    fmt.Println(utils.RedColor(err))
    return
  }

  fmt.Print("Podaj procent gęstości grafu: ")
  if _, err := fmt.Scanln(&density); err != nil {
    fmt.Println(utils.RedColor(err))
    return
  }

  edges, err := generator.GenerateGraph(numberOfVertices, 20, density, true)
  if err != nil {
    fmt.Println(utils.RedColor(err))
    return
  }

  mm.Analyzer, err = analysis.NewGraphAnalyzer(numberOfVertices, edges)
  if err != nil {
    fmt.Println(utils.RedColor(err))
    return
  }

  fmt.Println(utils.GreenColor("Wygenerowano graf"))

  var save string
  for {
    fmt.Print("Czy zapisać graf do pliku? [t/n] ")
    if _, err = fmt.Scanln(&save); err == nil {
      switch save {
      case "n":
        return
      case "t":
        var filename string
        fmt.Print("Podaj nazwe pliku: ")
        if _, err := fmt.Scanln(&filename); err != nil {
          fmt.Println(utils.RedColor(err))
          return
        }

        if err = utils.SaveGraph(filename, mm.Analyzer.List); err != nil {
          fmt.Println(utils.RedColor(err))
          return
        }

        fmt.Println(utils.GreenColor("Zapisano graf do pliku"))
        return
      default:
        fmt.Println("Wpisz t lub n")
      }
    } else {
      fmt.Println(utils.RedColor(err))
    }
  }
}

func (mm *MenuManager) readGraphFromFile() {
  fmt.Print("Podaj ścieżekę do pliku: ")
  var filename string
  if _, err := fmt.Scanln(&filename); err != nil {
    fmt.Println(utils.RedColor(err))
    return
  }

  numberOfVertices, edges, err := utils.ReadGraphFromFile(filename)
  if err != nil {
    fmt.Println(utils.RedColor(err))
    return
  }

  mm.Analyzer, err = analysis.NewGraphAnalyzer(numberOfVertices, edges)
  if err != nil {
    fmt.Println(utils.RedColor(err))
    return
  }
  fmt.Println(utils.GreenColor("Wczytano graf z pliku"))
}

func (mm *MenuManager) chooseProblem() {
  for {
    fmt.Println("Wybierz problem do rozwiązania:")
    fmt.Println("0. Wyjście")
    fmt.Println("1. Wyznaczanie minimalnego drzewa rozpinającego")
    fmt.Println("2. Wyznaczanie najkrótszej ścieżki w grafie")
    fmt.Print("Twój wybór: ")

    var choice int
    if _, err := fmt.Scanln(&choice); err != nil {
      fmt.Println(utils.RedColor(err))
      continue
    }

    switch choice {
    case 0:
      return
    case 1:
      mm.chooseAlgorithm(mstProblem)
    case 2:
      mm.chooseAlgorithm(shortestPathProblem)
    default:
      fmt.Println("Tylko opcje 0-2")
    }
  }
}

func (mm *MenuManager) chooseAlgorithm(problem int) {
  if mm.Analyzer == nil {
    fmt.Println("Brak grafu do przeanalizowania")
    return
  }

  switch problem {
  case mstProblem:
    for {
      fmt.Println("Wybierz algorytm:")
      fmt.Println("0. Wyjście")
      fmt.Println("1. Prim")
      fmt.Println("2. Kruskal")
      fmt.Print("Twój wybór: ")
      var choice int
      if _, err := fmt.Scanln(&choice); err != nil {
        fmt.Println(utils.RedColor(err))
        continue
      }

      switch choice {
      case 0:
        return
      case 1:
        mm.Analyzer.Prim(true)
      case 2:
        mm.Analyzer.Kruskal(true)
      default:
        fmt.Println("Tylko opcje 0-2")
      }
    }
  case shortestPathProblem:
    for {
      fmt.Println("Wybierz algorytm:")
      fmt.Println("0. Wyjście")
      fmt.Println("1. Dijsktra")
      fmt.Println("2. Ford-Bellman")
      fmt.Print("Twój wybór: ")
      var choice int
      if _, err := fmt.Scanln(&choice); err != nil {
        fmt.Println(utils.RedColor(err))
        continue
      }

      switch choice {
      case 0:
        return
      case 1:
        var begin, end int
        fmt.Print("Wierzchołek startowy: ")
        if _, err := fmt.Scanln(&begin); err != nil {
          fmt.Println(utils.RedColor(err))
          continue
        }

        fmt.Print("Wierzchołek końcowy: ")
        if _, err := fmt.Scanln(&end); err != nil {
          fmt.Println(utils.RedColor(err))
          continue
        }

        mm.Analyzer.Dijstkra(begin, end)

      case 2:
        var begin, end int
        fmt.Print("Wierzchołek startowy: ")
        if _, err := fmt.Scanln(&begin); err != nil {
          fmt.Println(utils.RedColor(err))
          continue
        }

        fmt.Print("Wierzchołek końcowy: ")
        if _, err := fmt.Scanln(&end); err != nil {
          fmt.Println(utils.RedColor(err))
          continue
        }

        mm.Analyzer.BellmanFord(begin, end)
      default:
        fmt.Println("Tylko opcje 0-2")
      }
    }
  default:
    fmt.Println("Tylko opcje 0-2")
  }
}

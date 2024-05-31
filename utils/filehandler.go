package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Ite-2022-pwr/sem4-aizo-proj2-ak/graph"
)

// SaveCSV zapisuje dane do pliku csv
func SaveCSV(filename string, data [][]string) {
  fh, err := os.Create(filename)
  defer fh.Close()
  if err != nil {
    log.Fatal(RedColor("[!!] Nie udało się utworzyć pliku: ", filename))
  }

  wrtr := csv.NewWriter(fh)

  if err = wrtr.WriteAll(data); err != nil {
    log.Fatal(RedColor("[!!] Nie udało się zapisać danych do pliku: ", filename))
  }

  log.Println(GreenColor("[+] Zapisano dane do pliku: ", filename))
}

func ReadGraphFromFile(filename string) (numberOfVertices int, edges []graph.Edge, err error) {
  log.Println(YellowColor("[*] Czytanie danych z pliku: ", filename))
  f, err := os.Open(filename)
  if err != nil {
    log.Println(RedColor(fmt.Sprintf("[!!] Nie można otworzyć pliku: %v", err)))
    return 0, nil, err
  }
  defer f.Close()

  rdr := bufio.NewReader(f)
  var line string
  line, err = rdr.ReadString('\n')
  if err != nil {
    log.Println(RedColor(fmt.Sprintf("[!!] Błąd czytania linii '%v': %v", line, err)))
    return 0, nil, err
  }
  line = strings.TrimSpace(line)

  graphDataStr := strings.Fields(line)
  if len(graphDataStr) != 2 {
    log.Println(RedColor("[!!] Pierwsza linia pliku powinna zawierać liczbę krawędzi i wierzchołków grafu rozdzielone spacją"))
    return 0, nil, fmt.Errorf("[!!] Pierwsza linia pliku powinna zawierać liczbę krawędzi i wierzchołków grafu rozdzielone spacją")
  }

  numberOfEdges, err := strconv.Atoi(graphDataStr[0])
  if err != nil {
    log.Println(RedColor(err))
    return 0, nil, err
  }

  numberOfVertices, err = strconv.Atoi(graphDataStr[1])
  if err != nil {
    log.Println(RedColor(err))
    return 0, nil, err
  }

  if numberOfEdges < 1 || numberOfVertices < 1 {
    log.Println(RedColor("[!!] Liczba krawędzi oraz wierzchołków musi być dodatnia"))
    return 0, nil, fmt.Errorf("[!!] Liczba krawędzi oraz wierzchołków musi być dodatnia")
  }

  for i := 0; i < numberOfEdges; i++ {
    line, err = rdr.ReadString('\n')
    if err != nil {
      log.Println(RedColor(fmt.Sprintf("[!!] Błąd czytania linii '%v': %v", line, err)))
      return
    }

    edgeDataStr := strings.Fields(strings.TrimSpace(line))
    if len(edgeDataStr) != 3 {
      log.Println(RedColor(fmt.Sprintf("[!!] Linia %v: każda linia oprócz pierwszej powinna zawierać wierzchołek początkowy, końcowy oraz wagę krawędzi rozdzielone spacją", i + 1)))
      return 0, nil, fmt.Errorf("[!!] Linia %v: każda linia oprócz pierwszej powinna zawierać wierzchołek początkowy, końcowy oraz wagę krawędzi rozdzielone spacją", i + 1)
    }

    begin, err := strconv.Atoi(edgeDataStr[0])
    if err != nil {
      log.Println(RedColor(err))
      return 0, nil, err
    }

    end, err := strconv.Atoi(edgeDataStr[1])
    if err != nil {
      log.Println(RedColor(err))
      return 0, nil, err
    }

    weight, err := strconv.Atoi(edgeDataStr[2])
    if err != nil {
      log.Println(RedColor(err))
      return 0, nil, err
    }

    edges = append(edges, graph.Edge{Source: begin, Destination: end, Weight: weight})
  }

  return numberOfVertices, edges, nil
}

func SaveGraph(filename string, G graph.Graph) error {
  f, err := os.Create(filename)
  if err != nil {
    log.Println(RedColor(fmt.Sprintf("[!!] Nie można utworzyć pliku: %v", err)))
    return err
  }
  defer f.Close()

  wrtr := bufio.NewWriter(f)

  edges := G.GetEdges()
  n := len(edges)
  line := fmt.Sprintf("%v %v\n", G.GetEdgesNumber(), G.GetVerticesNumber())
  _, err = wrtr.WriteString(line)
  if err != nil {
    log.Println(RedColor(fmt.Sprintf("[!!] Nie można zapisać do pliku: %v", err)))
    return err
  }

  for i := 0; i < n; i++ {
    line = fmt.Sprintf("%v %v %v\n", edges[i].Source, edges[i].Destination, edges[i].Weight)
    _, err = wrtr.WriteString(line)
    if err != nil {
      log.Println(RedColor(fmt.Sprintf("[!!] Nie można zapisać do pliku: %v", err)))
      return err
    }
    wrtr.Flush()
  }

  log.Println(GreenColor(fmt.Sprintf("[+] Zapisano graf do pliku %v", filename)))

  return nil
}

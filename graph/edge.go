package graph

type Edge struct {
  Source      int
  Destination int
  Weight      int
}


func CompareEdges(e1, e2 Edge) int {
  if e1.Weight > e2.Weight {
    return 1
  } else if e1.Weight < e2.Weight {
    return -1
  }
  return 0
}

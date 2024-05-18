package graph

type Graph interface {
  // AddVertices(n int) error
  // AddEdge(u, v, cost int) error
  ToString() string
  LookupEdge(source, destination int) (bool, Edge, error)
  GetNeighbors(vertex int) ([]Edge, error)
}

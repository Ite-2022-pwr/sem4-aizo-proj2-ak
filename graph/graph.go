package graph

type Graph interface {
  // AddVertices(n int) error
  // AddEdge(u, v, cost int) error
  GetVerticesNumber() int
  GetEdgesNumber() int
  ToString() string
  LookupEdge(source, destination int, isDirected bool) (bool, Edge, error)
  GetNeighbors(vertex int) ([]Edge, error)
  GetEdges() []Edge
}

package graph

type Graph interface {
  AddVertices(n int) error
  AddEdge(u, v, cost int) error
  ToString() string
}

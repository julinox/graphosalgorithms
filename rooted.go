/* Package */
package graphosalgorithms

/* Imports */
import (
  "github.com/julinox/go_data_structures/graphos"
)

/* Glocals */

/* Types */

/* Interface */

/* Functions */
func RootTree(graph graphos.Grapho, node int) (graphos.Grapho) {

  var visited map[int]bool

  if (graph == nil) {
    return nil
  }

  rootGraph := graphos.InitGraphList()
  rootGraph.Flags |= graphos.GRAPH_DIRECTED
  visited = make(map[int]bool)
  rootTree(graph, rootGraph, node, visited)
  return rootGraph
}

func rootTree(graph graphos.Grapho, rootGraph graphos.Grapho, node int, visited map[int]bool) {

  /*
    Given a node make it root (careful of which you pick)

    - Create a new tree (being build at 'DFS' time)
    - Do not modify original tree
  */

  if (graph == nil) {
    return
  }

  neighbours := graph.VertexNeighbours(node)
  visited[node] = true
  for _, v := range *neighbours {
    if !(visited[v]) {
      rootGraph.EdgeAdd(node, v, graph.EdgeWeight(node, v))
      rootTree(graph, rootGraph, v, visited)
    }
  }
}

/*
  Given a rooted tree, set each node height. Its assumed that each node
  has a 'VertexInfo' structure asociated
*/

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
func NodeHeight(graph graphos.Grapho, node int) (int) {

  /*
    Get node Height
  */

  var h int

  if (graph == nil) {
    return -1
  }

  neighbours := graph.VertexNeighbours(node)
  if (neighbours == nil || len(*neighbours) == 0) {
    return 0
  }

  for _, v := range *neighbours {
    aux := NodeHeight(graph, v) + 1
    if (aux > h) {
      h = aux
    }
  }

  return h
}

func NodeDepth(graph graphos.Grapho, node int, height int) {

  /*
    Set the depth (heigh) for every node. Graph MUST be a tree.
    First call must start with the tree's root. If not then
    expect inaccurate outcome
  */

  if (graph == nil) {
    return
  }

  // Set my height
  graph.VertexAddDepth(node, height)
  neighbours := graph.VertexNeighbours(node)
  if (neighbours == nil) {
    return
  }

  for _, v := range *neighbours {
    NodeDepth(graph, v, height - 1)
  }
}

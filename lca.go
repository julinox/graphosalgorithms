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
func LowestCommonAncestor(graph graphos.Grapho, root, v1, v2 int) (int) {

  /*
    Find lowest common ancestor between v1 and v2
    Eulerian path method:
    - Tracking nodes path
    - Tracking nodes path heigh
    - Tracking nodes last 'path position' index

    Warning:
    - graph must be a rooted-tree
    - 0 <= node_id < number_of_nodes (for every node_id)

    Since node id's are expected to be positive a negative
    return value means some type of error
  */

  var vLen int
  var nodes []int
  var depth []int
  var vertices *[]int
  var lastPosition []int
  var li, lj, min, minPos int

  if (graph == nil) {
    return -1
  }

  if !(graph.GraphFlags() & graphos.GRAPH_DIRECTED == graphos.GRAPH_DIRECTED) {
    return -2
  }

  // This should not be done since its time consuming but safety first
  vertices = graph.VertexList()
  vLen = len(*vertices)
  lastPosition = make([]int, vLen)

  if (v1 >= vLen || v2 >= vLen) {
    return -3
  }

  for i := 0; i < vLen; i++ {
    if ((*vertices)[i] >= vLen) {
      return -4
    }

    lastPosition[i] = -1
  }

  lca(graph, root, &nodes, &depth, &lastPosition, 0)
  if (lastPosition[v1] > lastPosition[v2]) {
    li = lastPosition[v2]
    lj = lastPosition[v1]

  } else {
    li = lastPosition[v1]
    lj = lastPosition[v2]
  }

  min = depth[li]
  minPos = li
  for i := li; i <= lj; i++ {
    if (min > depth[i]) {
      min = depth[i]
      minPos = i
    }
  }

  return nodes[minPos]
}

func lca(graph graphos.Grapho, v int, nodes, depth, lastP *[]int, cHeight int) {

  if (graph == nil || nodes == nil) {
    return
  }

  for _, n := range *graph.VertexNeighbours(v) {
    *nodes = append(*nodes, v)
    *depth = append(*depth, cHeight)
    (*lastP)[v] = len(*nodes) - 1
    lca(graph, n, nodes, depth, lastP, cHeight + 1)
  }

  *nodes = append(*nodes, v)
  *depth = append(*depth, cHeight)
  (*lastP)[v] = len(*nodes) - 1
}

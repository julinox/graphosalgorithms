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
func CenterUndirected(graph graphos.Grapho) ([]int) {

  /*
    Get center node of graph
  */

  var target []int
  var ret []int
  var vxsDegree map[int]int

  if (graph == nil) {
    return []int{}
  }

  if (graph.GraphFlags() & graphos.GRAPH_DIRECTED == graphos.GRAPH_DIRECTED) {
    return []int{}
  }

  ret = make([]int, 0)
  target = make([]int, 0)
  vxsDegree = make(map[int]int)

  // Set starting degrees
  for _, v := range *graph.VertexList() {
    vxsDegree[v] = len(*graph.VertexNeighbours(v))
  }

  for {
    if (len(vxsDegree) <= 2) {
      break
    }

    // Checking for leaf nodes (target for 'elimination')
    for v, d := range vxsDegree {
      if (d <= 1) {
        target = append(target, v)
      }
    }

    for _, v := range target {
      neighbours := *graph.VertexNeighbours(v)
      if (len(neighbours) < 1) {
        continue
      }

      for _, n := range neighbours {
        if (vxsDegree[n] == 0) {
          continue
        }
        vxsDegree[n] -= 1
      }

      delete(vxsDegree, v)
    }

    target = nil
  }

  for v, _ := range vxsDegree {
    ret = append(ret, v)
  }

  return ret
}

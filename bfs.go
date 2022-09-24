/*
  ###############################################################################
  #                                                                             #
  # Breadth first search and all algorithms that use bfs                        #
  # This works only on graph which vertex's ID is 0 <= VID < |Vertices|         #
  #                                                                             #
  # BFS is a truly interesting algorithm since it allows to do things like:     #
  # - Do a graph traversal                                                      #
  # - Check for cycles                                                          #
  # - Check for connectivy between nodes                                        #
  # - know the shortest path (in terms of vertex distance) between one vertex   #
  #   and the rest                                                              #
  ###############################################################################
*/

/* Package */
package graphosalgorithms

/* Imports */
import (
  "github.com/julinox/go_data_structures/queue"
  "github.com/julinox/go_data_structures/graphos"
)

/* Glocals */

/* Types */
type BfsReturn struct {
  Path *[]int
  Distance *[]int
  Order *[]interface{}
}

/* Interface */

/* Functions */
func Bfs(graph graphos.Grapho, s, e int) (*BfsReturn) {

  /*
    s: Start vertex
    e: End vertex (optional)
    Path: Track the vertex where you came from
    Distance: Track the number of jumps (or layers) needed to reach
    the 'current' vertex
    qq: Queue of neighbours (bfs traversal)

    If 'e' < 0 means not to stop at 'e' but to keep doing the traversal
  */

  var forceStop bool
  var ret BfsReturn
  var vxtLen int
  var path []int
  var distance []int
  var visited []bool
  var qq *queue.Queue

  if (graph == nil || s < 0) {
    return nil
  }

  // Init
  forceStop = false
  vxtLen = graph.VertexMax()
  qq = queue.InitQueue(uint(vxtLen + 1))
  path = make([]int, vxtLen + 1)
  distance = make([]int, vxtLen + 1)
  visited = make([]bool, vxtLen + 1)
  for i := range path {
    path[i] = -1
    distance[i] = -1
  }

  // OPS
  qq.Enqueue(s)
  visited[s] = true
  distance[s] = 0
  for !qq.IsEmpty() && !forceStop {
    current, ta_ok := qq.Dequeue().(int)
    if (!ta_ok) { continue }

    neighbours := graph.VertexNeighbours(current)
    for _, next := range *neighbours {
      if !(visited[next]) {
        qq.Enqueue(next)
        visited[next] = true
        path[next] = current
        distance[next] = distance[current] + 1
        if (next == e) {
          forceStop = true
        }
      }
    }
  }

  ret.Path = &path
  ret.Distance = &distance
  ret.Order = qq.GetQueue()
  return &ret
}

func BfsShortestPathVector(graph graphos.Grapho, s int) (*[]int) {

  /*
    Get the shortest path vector between 's' and the rest
  */

  var bfsRet *BfsReturn

  if (graph == nil || s < 0) {
    return &[]int{}
  }

  bfsRet = Bfs(graph, s, -1)
  return bfsRet.Distance
}

func BfsShortestPath(graph graphos.Grapho, s, e int) (*[]int) {

  /*
    Get the shortest path between 's' and 'e'
  */

  var bfsRet *BfsReturn

  if (graph == nil || s < 0 || e < 0) {
   return &[]int{}
  }

  bfsRet = Bfs(graph, s, e)
  return ReconstructPath(bfsRet.Path, s, e)
}

func BfsShortestPathAll(graph graphos.Grapho, s int) (map[int]*[]int) {

  /*
    Returns the shortest path between 's' and the rest of vertices.
  */

  var pr *[]int
  var ret map[int]*[]int

  if (graph == nil || s < 0) {
    return nil
  }

  ret = make(map[int]*[]int)
  for _, v := range *graph.VertexList() {
    pr = BfsShortestPath(graph, s, v)
    ret[v] = pr
  }

  return ret
}

func ReconstructPath(rPath *[]int, s, e int) (*[]int) {

  var at int
  var path []int
  var reversed []int

  if (rPath == nil || e < 0) {
    return &[]int{}
  }

  at = e
  for at >= 0 {
    path = append(path, at)
    at = (*rPath)[at]
  }

  for i := len(path) -1; i >= 0; i-- {
    reversed = append(reversed, path[i])
  }

  if (reversed[0] != s) {
    return &[]int{}
  }

  return &reversed
}

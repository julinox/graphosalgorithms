/* Package */
package graphosalgorithms

/* Imports */
import (
  "github.com/julinox/go_data_structures/queue"
  "github.com/julinox/go_data_structures/graphos"
)

/* Glocals */

/* Types */
type ViStack struct {
  Stacked bool
  Visited bool
}

/* Interface */

/* Functions */
func TopSort(graph graphos.Grapho) (*[]int) {

  /*
    Good old topological sort.
    - Only DAG graphs have topsort ordering
    - Never returns nil
  */

  var order []int
  var orderPos int
  var vxs map[int]ViStack

  if (graph == nil) {
    return &([]int{})
  }

  vertices := graph.VertexList()
  orderPos = len(*vertices) - 1
  vxs = make(map[int]ViStack)
  order = make([]int, orderPos + 1)
  for _, v := range *vertices {
    if !(vxs[v].Visited) {
      topSort(graph, v, &order, &orderPos, vxs)
    }
    if (graph.GraphFlags() & graphos.GRAPH_HAS_CYCLE == graphos.GRAPH_HAS_CYCLE) {
      return &([]int{})
    }
  }

  return &order
}

func topSort(graph graphos.Grapho, v int, order *[]int, orderPos *int, vxs map[int]ViStack) {

  if (graph == nil || order == nil) {
    return
  }

  if (vxs[v].Visited) {
    return
  }

  if (vxs[v].Stacked) {
    graph.GraphFlagSet(graphos.GRAPH_HAS_CYCLE)
    return
  }


  vxs[v] = ViStack{true, false}
  neighbours := graph.VertexNeighbours(v)
  if (len(*neighbours) <= 0) {
    (*order)[*orderPos] = v
    *orderPos = *orderPos - 1
    vxs[v] = ViStack{true, true}
    return
  }

  for _, n := range *neighbours {
    topSort(graph, n, order, orderPos, vxs)
  }
  (*order)[*orderPos] = v
  *orderPos = *orderPos - 1
  vxs[v] = ViStack{true, true}
}

func TopSortKahn(graph graphos.Grapho) (*[]int) {

  /*
    Kahn's method (All vertices with input_degree = 0 are 'executable')
    - Enqueue all vertices with input_degree = 0
    - Remove enqueued vertices (append to order)
    - Recalculate input_degree
    - Repeat
  */

  var order []int
  var stack *queue.Queue
  var vDegree map[int]int

  if (graph == nil) {
    return &[]int{}
  }

  vertices := *graph.VertexList()
  order = make([]int, 0)
  vDegree = make(map[int]int)
  stack = queue.InitQueue(uint(len(vertices) + 1))

  // Set initial degrees
  for _, v := range vertices {
    vDegree[v] = graph.VertexInDegree(v)
    if (vDegree[v] == 0) {
      stack.Enqueue(v)
    }
  }

  for !stack.IsEmpty() {
    v, ok := stack.Dequeue().(int)
    if !(ok) {
      continue
    }

    neighbours := graph.VertexNeighbours(v)
    for _, n := range *neighbours {
      vDegree[n] -= 1
      if (vDegree[n] <= 0) {
        stack.Enqueue(n)
      }
    }

    delete(vDegree, v)
    order = append(order, v)
  }

  return &order
}

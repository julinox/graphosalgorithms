/* Package */
package graphosalgorithms

/* Imports */
import (
  "fmt"
  "sort"
  "github.com/julinox/go_data_structures/graphos"
)

/* Glocals */

/* Types */

/* Interface */

/* Functions */
func Ismorphic(g1, g2 graphos.Grapho) (bool) {

  /*
    Check wheter two graphs are ismorphic or not.
    Both g1 and g2 must be undirected graphs
  */

  if (g1 == nil || g2 == nil) {
    return false
  }

  if (g1.GraphFlags() & graphos.GRAPH_DIRECTED == graphos.GRAPH_DIRECTED) {
    return false
  }

  if (g2.GraphFlags() & graphos.GRAPH_DIRECTED == graphos.GRAPH_DIRECTED) {
    return false
  }

  tags1 := GraphEncode(g1)
  tags2 := GraphEncode(g2)

  for _, t1 := range tags1 {
    for _, t2 := range tags2 {
      if (t1 == t2) {
        return true
      }
    }
  }

  return false
}

func GraphEncode(graph graphos.Grapho) ([]string) {

  /*
    Encode an undirected graph:
    - Find a node candidate for be center
    - Create rooted-tree version from 'graph'
    - Get encode tags (a graph can have at most 2 center candidates)
  */

  var tags []string

  if (graph == nil) {
    return []string{}
  }

  if (graph.GraphFlags() & graphos.GRAPH_DIRECTED == graphos.GRAPH_DIRECTED) {
    return []string{}
  }

  center := CenterUndirected(graph)
  tags = make([]string, len(center))

  for i, _ := range center {
    tags[i] = TreeEncode(RootTree(graph, center[i]), center[i])
  }

  return tags
}

func TreeEncode(graph graphos.Grapho, vertex int) (string) {

  /*
    Tree graph encode: Uses AHU algorithm
  */

  var tag string
  var neighbourTags []string

  if (graph == nil) {
    return ""
  }

  tag = ""
  neighbours := *graph.VertexNeighbours(vertex)
  if (len(neighbours) <= 0) {
    return "()"
  }

  neighbourTags = make([]string, len(neighbours))
  for i, n := range neighbours {
    neighbourTags[i] = TreeEncode(graph, n)
  }

  sort.Strings(neighbourTags)
  for _, t := range neighbourTags {
    tag += t
  }

  return fmt.Sprintf("(%v)", tag)
}

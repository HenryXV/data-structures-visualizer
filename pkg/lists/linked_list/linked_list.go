package linked_list

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"log"
)

type List struct {
	dataFirst *node
	dataLast  *node
	size      int
}

type node struct {
	data interface{}

	prev *node
	next *node
}

func New() *List {
	return &List{}
}

func (l *List) Add(data interface{}) {
	newNode := &node{data: data, prev: l.dataLast}

	if l.size == 0 {
		l.dataFirst = newNode
		l.dataLast = newNode
	} else {
		l.dataLast.next = newNode
		l.dataLast = newNode
	}

	l.size += 1
}

func (l *List) Remove(index int) {
	if !l.isIndexValid(index) {
		return
	}

	var node *node
	if l.size-index < index {
		node = l.dataLast
		for n := l.size - 1; n != index; n, node = n-1, node.prev {
		}
	} else {
		node = l.dataFirst
		for n := 0; n != index; n, node = n+1, node.next {
		}
	}

	l.size -= 1

	if node == l.dataFirst {
		l.dataFirst = l.dataFirst.next
		node = nil
		return
	}

	if node == l.dataLast {
		l.dataLast = l.dataLast.prev
		node = nil
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	node = nil
}

func (l *List) Get(index int) (interface{}, bool) {
	if !l.isIndexValid(index) {
		return nil, false
	}

	if l.size-index < index {
		node := l.dataLast
		for n := l.size - 1; n != index; n, node = n-1, node.prev {
		}

		return node.data, true
	}

	node := l.dataFirst
	for n := 0; n != index; n, node = n+1, node.next {
	}

	return node.data, true
}

func (l *List) Contains(data interface{}) bool {

	if l.size == 0 {
		return false
	}

	node := l.dataFirst

	for {
		if node.data == data {
			return true
		}

		if node.next == nil {
			break
		}

		node = node.next
	}

	return false
}

func (l *List) Swap(i, j int) {
	if !l.isIndexValid(i) || !l.isIndexValid(j) {
		return
	}

	if i != j {
		var nodeI, nodeJ *node

		node := l.dataFirst
		for n := 0; nodeI == nil || nodeJ == nil; n, node = n+1, node.next {
			switch n {
			case i:
				nodeI = node
			case j:
				nodeJ = node
			}
		}

		nodeI.data, nodeJ.data = nodeJ.data, nodeI.data
	}
}

func (l *List) Insert(index int, data interface{}) {
	if !l.isIndexValid(index) {
		if index == l.size {
			l.Add(data)
		}
		return
	}

	var nodeAtIndex *node
	newNode := &node{data: data}

	if l.size-index < index {
		nodeAtIndex = l.dataLast
		for n := l.size - 1; n != index; n, nodeAtIndex = n-1, nodeAtIndex.prev {
		}
	} else {
		nodeAtIndex = l.dataFirst
		for n := 0; n != index; n, nodeAtIndex = n+1, nodeAtIndex.next {
		}
	}

	l.size += 1

	if nodeAtIndex == l.dataFirst {
		l.dataFirst = newNode
		l.dataFirst.next = nodeAtIndex
		nodeAtIndex.prev = newNode
		return
	}

	newNode.next = nodeAtIndex
	newNode.prev = nodeAtIndex.prev
	nodeAtIndex.prev.next = newNode
	nodeAtIndex.prev = newNode

}

// Visualizer TODO: fix when nodes are the same, do not circle back to a preexisting node
func (l *List) Visualizer() bytes.Buffer {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()

	l.createNodesAndEdges(l.dataFirst, graph, nil)

	var buf bytes.Buffer
	if err := g.Render(graph, graphviz.PNG, &buf); err != nil {
		log.Fatal(err)
	}

	return buf
}

func (l *List) createNodesAndEdges(node *node, graph *cgraph.Graph, prevGraphNode *cgraph.Node) {
	if node == nil {
		return
	}

	n, err := graph.CreateNode(fmt.Sprintf("%v", node.data))
	if err != nil {
		log.Fatal(err)
	}

	if prevGraphNode != nil {
		_, err = graph.CreateEdge("e", prevGraphNode, n)
		if err != nil {
			log.Fatal(err)
		}
	}

	l.createNodesAndEdges(node.next, graph, n)
}

func (l *List) Size() int {
	return l.size
}

func (l *List) isIndexValid(index int) bool {
	return index >= 0 && index < l.size
}

package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Tree(t *testing.T) {
	tree := NewTree()

	nodeA := &Node{
		Id:     "A",
		Name:   "Node A",
		Points: 0,
	}

	nodeA1 := &Node{
		Id:       "A1",
		Name:     "Node A1",
		ParentId: &nodeA.Id,
		Points:   4,
	}

	nodeA2 := &Node{
		Id:       "A2",
		Name:     "Node A2",
		ParentId: &nodeA.Id,
		Points:   6,
	}

	nodeB := &Node{
		Id:     "B",
		Name:   "Node B",
		Points: 0,
	}

	nodeB1 := &Node{
		Id:       "B1",
		Name:     "Node B1",
		Points:   3,
		ParentId: &nodeB.Id,
	}

	nodeB2 := &Node{
		Id:       "B2",
		Name:     "Node B2",
		Points:   7,
		ParentId: &nodeB.Id,
	}

	nodeC := &Node{
		Id:     "C",
		Name:   "Node C",
		Points: 20,
	}

	nodes := []*Node{nodeA, nodeA1, nodeA2, nodeB, nodeB1, nodeB2, nodeC}
	tree.PopulateTree(nodes)
	assert.Equal(t, len(nodeA.Children), 2)
	assert.Equal(t, nodeA.Points, float64(10))
	assert.Equal(t, nodeB.Points, float64(10))
	assert.Equal(t, nodeA.PercentageWeight, float64(25))
	assert.Equal(t, nodeB.PercentageWeight, float64(25))
	assert.Equal(t, nodeC.PercentageWeight, float64(50))
	assert.Equal(t, nodeA1.PercentageWeight, float64(40))
	assert.Equal(t, nodeA2.PercentageWeight, float64(60))
	assert.Equal(t, nodeB1.PercentageWeight, float64(30))
	assert.Equal(t, nodeB2.PercentageWeight, float64(70))

}

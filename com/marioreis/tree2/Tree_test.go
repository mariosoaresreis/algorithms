package tree2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type NodeInstance struct {
	Id               string
	Name             string
	ParentId         *string
	Parent           *Node
	Children         []Node
	Points           float64
	WeightPercentage float64
	Order            int64
}

func (n *NodeInstance) GetId() string {
	return n.Id
}

func (n *NodeInstance) GetName() string {
	return n.Name
}

func (n *NodeInstance) GetParentId() *string {
	return n.ParentId
}

func (n *NodeInstance) GetParent() *Node {
	return n.Parent
}

func (n *NodeInstance) SetParent(parent Node) {
	n.Parent = &parent
}

func (n *NodeInstance) HasChildren() bool {
	return len(n.Children) > 0
}

func (n *NodeInstance) GetWeightPercentage() float64 {
	return n.WeightPercentage
}

func (n *NodeInstance) SetWeightPercentage(weightPercentage float64) {
	n.WeightPercentage = weightPercentage
}

func (n *NodeInstance) GetChildren() []Node {
	return n.Children
}

func (n *NodeInstance) AddChild(child Node) {
	n.Children = append(n.Children, child)
}

func (n *NodeInstance) GetPoints() float64 {
	return n.Points
}

func (n *NodeInstance) SetPoints(points float64) {
	n.Points = points
}

func (n *NodeInstance) GetOrder() int64 {
	return n.Order
}

func Test_Tree(t *testing.T) {
	tree := NewTree()

	nodeA := &NodeInstance{
		Id:     "A",
		Name:   "Node A",
		Points: 0,
	}

	nodeA1 := &NodeInstance{
		Id:       "A1",
		Name:     "Node A1",
		ParentId: &nodeA.Id,
		Points:   4,
	}

	nodeA2 := &NodeInstance{
		Id:       "A2",
		Name:     "Node A2",
		ParentId: &nodeA.Id,
		Points:   6,
	}

	nodeB := &NodeInstance{
		Id:     "B",
		Name:   "Node B",
		Points: 0,
	}

	nodeB1 := &NodeInstance{
		Id:       "B1",
		Name:     "Node B1",
		Points:   3,
		ParentId: &nodeB.Id,
	}

	nodeB2 := &NodeInstance{
		Id:       "B2",
		Name:     "Node B2",
		Points:   7,
		ParentId: &nodeB.Id,
	}

	nodeC := &NodeInstance{
		Id:     "C",
		Name:   "Node C",
		Points: 20,
	}
	//     TOTAL = 40 pts = 100%
	//                   GROUPS                                          UNGROUPED - 50% of total
	//     A = 10pts=25% of total;  B = 10pts=25% of total        C = 20pts=100% of ungrouped
	//   /   \                    /   \
	// A1     A2                 B1   B2
	// A1=4pts = 40% of A        B1=3pts = 30% of B
	// A2=6pts = 60% of A        B2=7pts = 70% of B
	nodes := []Node{nodeA, nodeA1, nodeA2, nodeB, nodeB1, nodeB2, nodeC}
	tree.PopulateTree(nodes)
	assert.Equal(t, len(nodeA.Children), 2)
	assert.Equal(t, nodeA.Points, float64(10))
	assert.Equal(t, nodeB.Points, float64(10))
	assert.Equal(t, nodeA.WeightPercentage, float64(25))
	assert.Equal(t, nodeB.WeightPercentage, float64(25))
	assert.Equal(t, nodeC.WeightPercentage, float64(100))
	assert.Equal(t, nodeA1.WeightPercentage, float64(40))
	assert.Equal(t, nodeA2.WeightPercentage, float64(60))
	assert.Equal(t, nodeB1.WeightPercentage, float64(30))
	assert.Equal(t, nodeB2.WeightPercentage, float64(70))

}

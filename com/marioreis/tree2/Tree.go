package tree2

type Node interface {
	GetId() string
	GetName() string
	GetParentId() *string
	GetParent() *Node
	SetParent(parent Node)
	GetChildren() []Node
	AddChild(child Node)
	GetPoints() float64
	SetPoints(points float64)
	GetWeightPercentage() float64
	SetWeightPercentage(weightPercentage float64)
	HasChildren() bool
}

type Tree struct {
	Nodes          map[string]Node // All nodes
	GroupedNodes   map[string]Node
	UngroupedNodes map[string]Node
	Total          float64
}

func NewTree() *Tree {
	tree := &Tree{
		Nodes:          make(map[string]Node),
		GroupedNodes:   make(map[string]Node),
		UngroupedNodes: make(map[string]Node),
	}

	return tree
}

func (t *Tree) PopulateTree(nodes []Node) {
	t.addNodes(nodes)
	// Parents and children
	t.buildRelationships(nodes)
	t.buildGroups(nodes)
	t.recalculatePercentages()
}

// addNodes stores nodes in an internal map
func (t *Tree) addNodes(nodes []Node) {
	for _, n := range nodes {
		t.Nodes[n.GetId()] = n
	}
}

// buildRelationships establish relationship between parents and children
func (t *Tree) buildRelationships(nodes []Node) {
	for _, n := range nodes {
		parentNode := t.FindParent(n)

		if parentNode != nil {
			t.AddChild(*parentNode, n)
		}
	}

}

// buildGroups - build Grouped Nodes: Root nodes with Children.
// Root nodes without children go to a special group called 'Ungrouped'
func (t *Tree) buildGroups(nodes []Node) {
	for _, n := range nodes {
		if n.GetParent() == nil {
			// Reset points, in case we already have points from database
			if len(t.GroupedNodes) == 0 && len(t.UngroupedNodes) == 0 {
				t.Total = 0
			}

			t.Total += n.GetPoints()

			if n.HasChildren() {
				t.GroupedNodes[n.GetId()] = n
			} else {
				// Isolated nodes (without parent and children) go to a group called 'ungrouped nodes'
				t.UngroupedNodes[n.GetId()] = n
			}

		}
	}
}

func (t *Tree) recalculatePercentages() {
	for _, node := range t.GroupedNodes {
		node.SetWeightPercentage(node.GetPoints() / t.Total * 100)
	}

	ungroupedTotal := float64(0)

	for _, node := range t.UngroupedNodes {
		ungroupedTotal += node.GetPoints()
	}

	for _, node := range t.UngroupedNodes {
		node.SetWeightPercentage(node.GetPoints() / ungroupedTotal * 100)
	}
}

// FindParent Finds parent node, if exists
func (t *Tree) FindParent(n Node) *Node {
	if n.GetParentId() == nil {
		return nil
	}

	node, ok := t.Nodes[*n.GetParentId()]

	if ok {
		return &node
	}

	return nil
}

func (t *Tree) AddChild(parent, child Node) {
	if parent != nil {
		// To reset any previous points from database
		if !parent.HasChildren() {
			parent.SetPoints(0)
		}

		parent.AddChild(child)
		parent.SetPoints(parent.GetPoints() + child.GetPoints())
		child.SetParent(parent)

		for _, n := range parent.GetChildren() {
			n.SetWeightPercentage(n.GetPoints() / parent.GetPoints() * 100)
		}
	}
}

package tree

type Node struct {
	Id               string
	Name             string
	ParentId         *string
	Parent           *Node
	Children         []*Node
	Points           float64
	PercentageWeight float64
}

func (n *Node) HasChildren() bool {
	return len(n.Children) > 0
}

type Tree struct {
	Nodes          map[string]*Node // All nodes
	GroupedNodes   map[string]*Node
	UngroupedNodes map[string]*Node
	Total          float64
}

func NewTree() *Tree {
	tree := &Tree{
		Nodes:          make(map[string]*Node),
		GroupedNodes:   make(map[string]*Node),
		UngroupedNodes: make(map[string]*Node),
	}

	return tree
}

func (t *Tree) PopulateTree(nodes []*Node) {
	t.addNodes(nodes)
	// Parents and children
	t.buildRelationships(nodes)
	t.buildGroups(nodes)
	t.recalculatePercentages()
}

// addNodes stores nodes in an internal map
func (t *Tree) addNodes(nodes []*Node) {
	for _, n := range nodes {
		t.Nodes[n.Id] = n
	}
}

// buildRelationships establish relationship between parents and children
func (t *Tree) buildRelationships(nodes []*Node) {
	for _, n := range nodes {
		parentNode := t.FindParent(n)

		if parentNode != nil {
			t.AddChild(parentNode, n)
		}
	}

}

// buildGroups - build grouped nodes - root nodes with children
// root nodes without children go to group named 'Ungrouped'
func (t *Tree) buildGroups(nodes []*Node) {
	for _, n := range nodes {
		if n.Parent == nil {
			// Reset points, in case we already have points from database
			if len(t.GroupedNodes) == 0 && len(t.UngroupedNodes) == 0 {
				t.Total = 0
			}

			t.Total += n.Points

			if n.HasChildren() {
				t.GroupedNodes[n.Id] = n
			} else {
				// Isolated nodes go to a group named 'ungrouped nodes'
				t.UngroupedNodes[n.Id] = n
			}

		}
	}
}

func (t *Tree) recalculatePercentages() {
	for _, node := range t.GroupedNodes {
		node.PercentageWeight = node.Points / t.Total * 100
	}

	for _, node := range t.UngroupedNodes {
		node.PercentageWeight = node.Points / t.Total * 100
	}

}

// FindParent Finds parent node, if exists
func (t *Tree) FindParent(n *Node) *Node {
	if n.ParentId == nil {
		return nil
	}

	return t.Nodes[*n.ParentId]
}

func (t *Tree) AddChild(parent, child *Node) {
	if parent != nil {
		// To reset any previous points from database
		if !parent.HasChildren() {
			parent.Points = 0
		}

		parent.Children = append(parent.Children, child)
		parent.Points += child.Points
		child.Parent = parent

		for _, n := range parent.Children {
			n.PercentageWeight = n.Points / parent.Points * 100
		}
	}
}

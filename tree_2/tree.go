package tree

import (
	"fmt"
	"sort"
)

type Records []Record

func BuildTree(records []Record) (*Node, error) {
	sort.Sort(Records(records))
	idToNode := make(map[int]*Node)

	for i, record := range records {
		if _,ok := idToNode[record.ID]; ok {
			return nil, fmt.Errorf("failed to build tree. Found duplicate node with id [%d]",record.ID)
		}
		if record.ID != i {
			if i == 0 {
				return nil,fmt.Errorf("failed to build tree. No root node given")
			}
			return nil, fmt.Errorf("failed to build tree. Non continous tree")
		}
		if record.ID == record.Parent && record.ID != 0 {
			return nil, fmt.Errorf("failed to build tree")
		}

		node := &Node{ID: record.ID}
		parentNode, ok := idToNode[record.Parent]

		if !ok {
			if record.ID == 0 && record.Parent > 0 {
				return nil,fmt.Errorf("failed to build tree. Root node has parent of id [%d]",record.Parent)
			}
			if record.ID != 0 {
				return nil,fmt.Errorf("failed to build tree, invalid parent")
			}
		} else {
			parentNode.Children = append(parentNode.Children, node)
		}
		idToNode[node.ID] = node
	}
	return idToNode[0], nil
}

func (r Records) Len() int {
	return len(r)
}

func (r Records) Less(i, j int) bool {
	return r[i].ID < r[j].ID
}

func (r Records) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

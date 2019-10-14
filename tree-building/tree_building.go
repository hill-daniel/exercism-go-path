// Package tree provides functionality to create a tree with records.
// The records only contain an ID number and a parent ID number. The ID number is always between
// 0 (inclusive) and the length of the record list (exclusive). All records have a parent ID lower
// than their own ID, except for the root record, which has a parent ID that's equal to its own ID.
package tree

import (
	"fmt"
	"sort"
)

// Record is used to construct a tree
type Record struct {
	ID, Parent int
}

// Node represents a node of the tree
type Node struct {
	ID       int
	Children []*Node
}

// Records is a custom type for enable sorting on a slice of Record.
type Records []Record

func (r Records) Len() int {
	return len(r)
}

func (r Records) Less(i, j int) bool {
	return r[i].ID < r[j].ID
}

func (r Records) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// Build creates a tree with a given slice of records.
func Build(records []Record) (*Node, error) {
	sort.Sort(Records(records))
	idToNode := make(map[int]*Node)

	for i, r := range records {
		if _, ok := idToNode[r.ID]; ok {
			return nil, fmt.Errorf("failed to build tree. Found duplicate node id [%d]", r.ID)
		}
		if r.ID != i {
			return nil, fmt.Errorf("failed to build tree. Non continous tree")
		}

		node := &Node{ID: r.ID}
		if parentNode, ok := idToNode[r.Parent]; ok {
			parentNode.Children = append(parentNode.Children, node)
		} else if !isRoot(r) {
			return nil, fmt.Errorf("failed to build tree, invalid parent [%d] of node [%s]", r.Parent, node)
		}
		idToNode[node.ID] = node
	}
	return idToNode[0], nil
}

func isRoot(record Record) bool {
	return record.ID == 0 && record.Parent == 0
}

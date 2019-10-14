package tree

import (
	"reflect"
	"testing"
)

func Test_should_create_one_node_tree(t *testing.T) {
	input := []Record{{ID: 0}}
	expected := &Node{
		ID: 0,
	}
	tree, err := BuildTree(input)
	if err != nil {
		t.Errorf("failed to build tree, %v", err)
	}
	if !reflect.DeepEqual(tree, expected) {
		t.Fatalf("returned %s but was expected to return %s.", tree, expected)
	}
}

func Test_should_create_tree_with_three_nodes_in_order(t *testing.T) {
	input := []Record{
		{ID: 0},
		{ID: 1, Parent: 0},
		{ID: 2, Parent: 0},
	}
	expected := &Node{
		ID: 0,
		Children: []*Node{
			{ID: 1},
			{ID: 2},
		}}
	tree, err := BuildTree(input)
	if err != nil {
		t.Errorf("failed to build tree, %v", err)
	}
	if !reflect.DeepEqual(tree, expected) {
		t.Fatalf("returned %s but was expected to return %s.", tree, expected)
	}
}

func Test_should_create_tree_with_three_nodes_in_reverse_order(t *testing.T) {
	input := []Record{
		{ID: 2},
		{ID: 1, Parent: 0},
		{ID: 0, Parent: 0},
	}
	expected := &Node{
		ID: 0,
		Children: []*Node{
			{ID: 1},
			{ID: 2},
		}}

	tree, err := BuildTree(input)

	if err != nil {
		t.Errorf("failed to build tree, %v", err)
	}
	if !reflect.DeepEqual(tree, expected) {
		t.Fatalf("returned %s but was expected to return %s.", tree, expected)
	}
}

func Test_should_create_tree_with_node_with_more_than_two_children(t *testing.T) {
	input := []Record{
		{ID: 3, Parent: 0},
		{ID: 2, Parent: 0},
		{ID: 1, Parent: 0},
		{ID: 0},
	}
	expected := &Node{
		ID: 0,
		Children: []*Node{
			{ID: 1},
			{ID: 2},
			{ID: 3},
		},
	}

	tree, err := BuildTree(input)

	if err != nil {
		t.Errorf("failed to build tree, %v", err)
	}
	if !reflect.DeepEqual(tree, expected) {
		t.Fatalf("returned %s but was expected to return %s.", tree, expected)
	}
}
func Test_should_create_binary_tree(t *testing.T) {
	input := []Record{
		{ID: 5, Parent: 1},
		{ID: 3, Parent: 2},
		{ID: 2, Parent: 0},
		{ID: 4, Parent: 1},
		{ID: 1, Parent: 0},
		{ID: 0},
		{ID: 6, Parent: 2},
	}
	expected := &Node{
		ID: 0,
		Children: []*Node{
			{
				ID: 1,
				Children: []*Node{
					{ID: 4},
					{ID: 5},
				},
			},
			{
				ID: 2,
				Children: []*Node{
					{ID: 3},
					{ID: 6},
				},
			},
		},
	}

	tree, err := BuildTree(input)

	if err != nil {
		t.Errorf("failed to build tree, %v", err)
	}
	if !reflect.DeepEqual(tree, expected) {
		t.Fatalf("returned %s but was expected to return %s.", tree, expected)
	}
}

func Test_should_create_unbalanced_tree(t *testing.T) {
	input := []Record{
		{ID: 5, Parent: 2},
		{ID: 3, Parent: 2},
		{ID: 2, Parent: 0},
		{ID: 4, Parent: 1},
		{ID: 1, Parent: 0},
		{ID: 0},
		{ID: 6, Parent: 2},
	}
	expected := &Node{
		ID: 0,
		Children: []*Node{
			{
				ID: 1,
				Children: []*Node{
					{ID: 4},
				},
			},
			{
				ID: 2,
				Children: []*Node{
					{ID: 3},
					{ID: 5},
					{ID: 6},
				},
			},
		},
	}
	tree, err := BuildTree(input)

	if err != nil {
		t.Errorf("failed to build tree, %v", err)
	}
	if !reflect.DeepEqual(tree, expected) {
		t.Fatalf("returned %s but was expected to return %s.", tree, expected)
	}
}

func Test_should_fail_if_root_node_has_parent(t *testing.T) {
	input := []Record{
		{ID: 0, Parent: 1},
		{ID: 1, Parent: 0},
	}

	_, err := BuildTree(input)

	if err == nil {
		t.Errorf("expected error, but none showed up")
	}
	if err.Error() != "failed to build tree. Root node has parent of id [1]" {
		t.Errorf("expected difertent error msg than %v", err)
	}
}

func Test_should_fail_if_no_root_is_present(t *testing.T) {
	input := []Record{
		{ID: 1, Parent: 0},
	}
	_, err := BuildTree(input)

	if err == nil {
		t.Errorf("expected error, but none showed up")
	}
	if err.Error() != "failed to build tree. No root node given" {
		t.Errorf("expected difertent error msg than %v", err)
	}
}
func Test_should_fail_if_node_is_duplicate(t *testing.T) {
	input := []Record{
		{ID: 0, Parent: 0},
		{ID: 1, Parent: 0},
		{ID: 1, Parent: 0},
	}
	_, err := BuildTree(input)

	if err == nil {
		t.Errorf("expected error, but none showed up")
	}
	if err.Error() != "failed to build tree. Found duplicate node with id [1]" {
		t.Errorf("expected difertent error msg than %v", err)
	}
}

func Test_should_fail_if_root_is_duplicate(t *testing.T) {
	input := []Record{
		{ID: 0, Parent: 0},
		{ID: 0, Parent: 0},
	}

	_, err := BuildTree(input)

	if err == nil {
		t.Errorf("expected error, but none showed up")
	}
	if err.Error() != "failed to build tree. Found duplicate node with id [0]" {
		t.Errorf("expected difertent error msg than %v", err)
	}
}

func Test_should_fail_if_tree_is_not_continuous(t *testing.T) {
	input := []Record{
		{ID: 2, Parent: 0},
		{ID: 4, Parent: 2},
		{ID: 1, Parent: 0},
		{ID: 0},
	}
	_, err := BuildTree(input)

	if err == nil {
		t.Errorf("expected error, but none showed up")
	}
	if err.Error() != "failed to build tree. Non continous tree" {
		t.Errorf("expected difertent error msg than %v", err)
	}
}
func Test_should_fail_if_nodes_form_direct_cycle(t *testing.T) {
	input := []Record{
		{ID: 5, Parent: 2},
		{ID: 3, Parent: 2},
		{ID: 2, Parent: 2},
		{ID: 4, Parent: 1},
		{ID: 1, Parent: 0},
		{ID: 0},
		{ID: 6, Parent: 3},
	}
	_, err := BuildTree(input)

	if err == nil {
		t.Errorf("expected error, but none showed up")
	}
	if err.Error() != "failed to build tree" {
		t.Errorf("expected difertent error msg than %v", err)
	}
}
func Test_should_fail_if_nodes_form_indirect_cycle(t *testing.T) {
	input := []Record{
		{ID: 5, Parent: 2},
		{ID: 3, Parent: 2},
		{ID: 2, Parent: 6},
		{ID: 4, Parent: 1},
		{ID: 1, Parent: 0},
		{ID: 0},
		{ID: 6, Parent: 3},
	}
	_, err := BuildTree(input)

	if err == nil {
		t.Errorf("expected error, but none showed up")
	}
	if err.Error() != "failed to build tree, invalid parent" {
		t.Errorf("expected difertent error msg than %v", err)
	}
}
func Test_should_fail_if_node_has_higher_id_parent_of_lower_id(t *testing.T) {
	input := []Record{
		{ID: 0},
		{ID: 2, Parent: 0},
		{ID: 1, Parent: 2},
	}
	_, err := BuildTree(input)

	if err == nil {
		t.Errorf("expected error, but none showed up")
	}
	if err.Error() != "failed to build tree, invalid parent" {
		t.Errorf("expected difertent error msg than %v", err)
	}
}

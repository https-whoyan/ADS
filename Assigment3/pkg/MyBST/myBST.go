package MyBST

import (
	"cmp"
	"errors"
	"fmt"
)

type ord cmp.Ordered

type node[K ord, V any] struct {
	key   K
	val   V
	left  *node[K, V]
	right *node[K, V]
}

type MyBST[K ord, V any] struct {
	root *node[K, V]
	keys map[K]bool
	size int
}

func NewMyBST[K ord, V any]() *MyBST[K, V] {
	return &MyBST[K, V]{
		root: nil,
		keys: make(map[K]bool),
	}
}

func (n *node[K, V]) appendNode(key K, value V) {
	nodeKey := n.key
	if key < nodeKey {
		if n.left == nil {
			n.left = &node[K, V]{
				key: key,
				val: value,
			}
			return
		}
		n.left.appendNode(key, value)
		return
	}
	if n.right == nil {
		n.right = &node[K, V]{
			key: key,
			val: value,
		}
		return
	}
	n.right.appendNode(key, value)
}

func (b *MyBST[K, V]) ExistKey(key K) bool {
	_, contains := b.keys[key]
	return contains
}

func (b *MyBST[K, V]) Put(key K, value V) error {
	if b.ExistKey(key) {
		return errors.New(fmt.Sprintf("Key %v is already exist!", key))
	}
	if b.size == 0 {
		b.root = &node[K, V]{
			key: key,
			val: value,
		}
		b.keys[key] = true
		b.size++
		return nil
	}
	b.keys[key] = true
	b.root.appendNode(key, value)
	b.size++
	return nil
}

func (n *node[K, V]) get(searchedKey K) (value V, err error) {
	if n == nil {
		return value, errors.New(fmt.Sprintf("Key %v not exist", searchedKey))
	}
	nodeKey := n.key
	if nodeKey == searchedKey {
		return n.val, nil
	}
	if searchedKey < nodeKey {
		return n.left.get(searchedKey)
	}
	return n.right.get(searchedKey)
}

func (b *MyBST[K, V]) Get(searchedKey K) (V, error) {
	return b.root.get(searchedKey)
}

func (n *node[K, V]) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func (b *MyBST[K, V]) Delete(key K) {
	// If key not contains, break
	if !b.ExistKey(key) {
		return
	}
	// Update size and keys
	b.size--
	delete(b.keys, key)

	// deleted node function
	var findDeletedNode func(curr *node[K, V], deletedKey K) *node[K, V]
	findDeletedNode = func(curr *node[K, V], deletedKey K) *node[K, V] {
		nodeKey := curr.key
		if nodeKey == deletedKey {
			return curr
		}
		// I already check, that deletedKey exist
		// That I'm sure to recursive call, I not get the nil pointer
		if deletedKey < nodeKey {
			return findDeletedNode(curr.left, deletedKey)
		}
		return findDeletedNode(curr.left, deletedKey)
	}
	// findParentOfDeletedNode function.
	// It is necessary to swap values of subtrees
	var findParentOfDeletedNode func(curr *node[K, V], searchedNode *node[K, V]) (parent *node[K, V], isLeft bool)
	findParentOfDeletedNode = func(curr *node[K, V], searchedNode *node[K, V]) (parent *node[K, V], isLeft bool) {
		// I'm comparing pointers, so it's okay.
		if curr.left == searchedNode {
			return curr, true
		} else if curr.right == searchedNode {
			return curr, false
		}

		// Recursive call:
		searchedNodeKey := searchedNode.key
		currNodeKey := curr.key
		if searchedNodeKey < currNodeKey {
			return findParentOfDeletedNode(curr.left, searchedNode)
		}
		return findParentOfDeletedNode(curr.right, searchedNode)
	}
	var findMinimumKey func(curr *node[K, V]) *node[K, V]
	findMinimumKey = func(curr *node[K, V]) *node[K, V] {
		if curr.isLeaf() {
			return curr
		}
		return findMinimumKey(curr.left)
	}
	// Get deletedNode pointer
	deletedNode := findDeletedNode(b.root, key)
	// I've a case, that deleted note is a root of BST
	// So, I check it:
	if b.root == deletedNode {
		// If a have only one root
		if b.root.isLeaf() {
			// So easy, just replace.
			b.root = nil
			return
		}
		// I've a case, when I haven't a right subtree.
		if b.root.right == nil {
			// In this case, I'm sure, that root left subtree exist, just
			// replace BST root to left subtree
			b.root = b.root.left
			return
		}
		// Else
		// I'll find a minimum key in right subtree
		// Delete a minimum node in right subtree (it will be a leaf)
		// And replace values.

		// It will be working because I know, that
		// Every key in right subtree will be large, than every key in left subtree
		rightSubtreeMinNode := findMinimumKey(b.root.right)
		rightSubtreeMinNodeParent, isLeft := findParentOfDeletedNode(b.root, rightSubtreeMinNode)
		// Replace parent of minimum key in right subtree node values
		if isLeft {
			rightSubtreeMinNodeParent.left = nil
		} else {
			rightSubtreeMinNodeParent.right = nil
		}
		// Ok, done
		// I know, that deletedNode is root of BST
		// Just replace his key and value
		b.root.key = rightSubtreeMinNode.key
		b.root.val = rightSubtreeMinNode.val
		// Done deleting, return
		return
	}
	// If I delete a leaf:
	deletedNodeParent, isLeft := findParentOfDeletedNode(b.root, deletedNode)
	if deletedNode.isLeaf() {
		// Just replace this node to nil
		if isLeft {
			deletedNodeParent.left = nil
		}
		deletedNodeParent.right = nil
		// Easy, return
		return
	}
	// I've a case that I haven't a right subtree
	// So easy, replace parent children to deleted node children
	if deletedNodeParent.right == nil {
		deletedNodeParent.left = deletedNode.left
		deletedNodeParent.right = deletedNode.right
	}
	// I'll find a minimum key in right subtree
	// Delete a minimum node in right subtree (it will be a leaf)
	// And replace values.

	// It will be working because I know, that
	// Every key in right subtree will be large, than every key in left subtree
	rightSubtreeMinNode := findMinimumKey(deletedNode.right)
	rightSubtreeMinNodeParent, isLeft := findParentOfDeletedNode(deletedNode, rightSubtreeMinNode)
	// Replace parent of minimum key in right subtree node values
	if isLeft {
		rightSubtreeMinNodeParent.left = nil
	} else {
		rightSubtreeMinNodeParent.right = nil
	}
	// Ok, done
	// Replace deleted node values
	deletedNode.key = rightSubtreeMinNode.key
	deletedNode.val = rightSubtreeMinNode.val
	// Done, return
}

type TraversalNode[K ord, V any] struct {
	key   K
	value V
}

func (n *node[K, V]) traversal() []TraversalNode[K, V] {
	// From my leetcode.com solution L)
	if n == nil {
		return []TraversalNode[K, V]{}
	}
	leftValues := n.left.traversal()
	currVal := []TraversalNode[K, V]{
		{
			key:   n.key,
			value: n.val,
		},
	}
	rightValues := n.right.traversal()
	return append(append(leftValues, currVal...), rightValues...)
}

func (b *MyBST[K, V]) InOrderTraversal() []TraversalNode[K, V] {
	return b.root.traversal()
}

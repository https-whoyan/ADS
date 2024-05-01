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
	size int
}

func NewMyBST[K ord, V any]() *MyBST[K, V] {
	return &MyBST[K, V]{
		root: nil,
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
	_, err := b.Get(key)
	return err == nil
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
		b.size++
		return nil
	}
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

func (n *node[K, V]) bfsPrint() {
	stack := []*node[K, V]{n}
	height := 1
	var nextStack []*node[K, V]
	for len(stack) != 0 {
		fmt.Printf("\nHeight: %v\n", height)
		for _, v := range stack {
			fmt.Print(v.key, ": left:")
			if v.left != nil {
				fmt.Print(v.left.key)
			}
			fmt.Print(", right: ")
			if v.right != nil {
				fmt.Print(v.right.key)
			}
			fmt.Print(";;;")
			if v.left != nil {
				nextStack = append(nextStack, v.left)
			}
			if v.right != nil {
				nextStack = append(nextStack, v.right)
			}
		}
		stack = nextStack
		nextStack = []*node[K, V]{}
		height++
	}
	fmt.Print("\n========\n")
}

func (b *MyBST[K, V]) Delete(key K) {
	//fmt.Printf("\n=======\nDelete %v", key)
	//b.root.bfsPrint()
	// If key not contains, break
	if !b.ExistKey(key) {
		return
	}
	// Update size
	b.size--

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
		return findDeletedNode(curr.right, deletedKey)
	}
	// findParentOfDeletedNode function.
	// It is necessary to swap values of subtrees
	var findParentOfDeletedNode func(curr *node[K, V], searchedNode *node[K, V]) (parent *node[K, V], isLeaf bool)
	findParentOfDeletedNode = func(curr *node[K, V], searchedNode *node[K, V]) (parent *node[K, V], isLeaf bool) {
		// I'm comparing pointers, so it's okay.
		if curr.left == searchedNode {
			return curr, searchedNode.isLeaf()
		} else if curr.right == searchedNode {
			return curr, searchedNode.isLeaf()
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
		if curr.left == nil {
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
		minNodeKey := rightSubtreeMinNode.key
		minNomeValue := rightSubtreeMinNode.val
		rightSubtreeMinNodeParent, isLeaf := findParentOfDeletedNode(b.root, rightSubtreeMinNode)
		// Replace parent of minimum key in right subtree node values
		//if minimum node is Leaf, stand parent node left tree nil
		if isLeaf {
			isLeft := rightSubtreeMinNodeParent.left == rightSubtreeMinNode
			if isLeft {
				rightSubtreeMinNodeParent.left = nil
			} else {
				rightSubtreeMinNodeParent.right = nil
			}
			// I know, that deletedNode is root of BST
			// Just replace his key and value
			b.root.key = minNodeKey
			b.root.val = minNomeValue
			return
		}
		// Else I stand, that rightSubtreeMinNode haven't a left subtree (rightSubtreeMinNode.left == nil).
		// But I have a right subtree
		// Ok, replace rightSubtreeMinNode node to rightSubtreeMinNode.right node
		// Then, replace parent minimum node right subtree to minNode subtree
		*rightSubtreeMinNode = *(rightSubtreeMinNode.right)
		// Ok, done
		// I know, that deletedNode is root of BST
		// Just replace his key and value
		b.root.key = minNodeKey
		b.root.val = minNomeValue
		// Done deleting, return
		return
	}
	deletedNodeParent, isLeaf := findParentOfDeletedNode(b.root, deletedNode)
	// If I delete a leaf:
	if isLeaf {
		// Just replace this node to nil
		isLeft := deletedNodeParent.left == deletedNode
		if isLeft {
			deletedNodeParent.left = nil
			return
		}
		deletedNodeParent.right = nil
		return
	}
	// I've a case that I haven't a right subtree

	// So easy, replace deleted node to deleted node left subtree root node
	// And, that's all

	// I can do this, because deletedNode is not a leaf.
	if deletedNode.right == nil {
		*(deletedNode) = *(deletedNode.left)
		return
	}
	// I'll find a minimum key in right subtree
	// Delete a minimum node in right subtree (it will be a leaf)
	// And replace values.

	// It will be working because I know, that
	// Every key in right subtree will be large, than every key in left subtree
	rightSubtreeMinNode := findMinimumKey(deletedNode.right)
	minNodeKey := rightSubtreeMinNode.key
	minNomeValue := rightSubtreeMinNode.val
	rightSubtreeMinNodeParent, isLeaf := findParentOfDeletedNode(deletedNode, rightSubtreeMinNode)
	// Replace parent of minimum key in right subtree node values
	//if minimum node is Leaf, stand parent node left tree nil
	if isLeaf {
		isLeft := rightSubtreeMinNodeParent.left == rightSubtreeMinNode
		if isLeft {
			rightSubtreeMinNodeParent.left = nil
		} else {
			rightSubtreeMinNodeParent.right = nil
		}
		// Just replace his key and value
		deletedNode.key = minNodeKey
		deletedNode.val = minNomeValue
		// Done, return
		return
	}
	// Else I stand, that rightSubtreeMinNode haven't a left subtree (rightSubtreeMinNode.left == nil).
	// Then, replace parent minimum node right subtree to minNode subtree
	rightSubtreeMinNodeParent.right = rightSubtreeMinNode.right
	// Ok, done
	// I know, that deletedNode is root of BST
	// Just replace his key and value
	deletedNode.key = minNodeKey
	deletedNode.val = minNomeValue
	// Done deleting, return
	return
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

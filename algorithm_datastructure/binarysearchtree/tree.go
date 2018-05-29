// Package binarysearchtree 非平衡二分搜索树
// 	From https://flaviocopes.com/golang-data-structure-binary-search-tree/
package binarysearchtree

import (
	"fmt"
)

// Node 节点
type Node struct {
	key   string
	value interface{}
	left  *Node
	right *Node
}

// BinarySearchTree 二分搜索树
type BinarySearchTree struct {
	root *Node
}

// Insert 添加元素
func (bst *BinarySearchTree) Insert(key string, value interface{}) {
	node := &Node{key, value, nil, nil}
	if bst.root == nil {
		bst.root = node
		return
	}
	insertNode(bst.root, node)
}

func insertNode(node, newNode *Node) {
	if node.key > newNode.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

// InOrderTraverse 遍历 -- 左中右
func (bst *BinarySearchTree) InOrderTraverse(f func(interface{})) {
	inOrderTraverse(bst.root, f)
}

func inOrderTraverse(node *Node, f func(interface{})) {
	if node != nil {
		inOrderTraverse(node.left, f)
		f(node.value)
		inOrderTraverse(node.right, f)
	}
}

// PreOrderTraverse 遍历 -- 中左右
func (bst *BinarySearchTree) PreOrderTraverse(f func(interface{})) {
	preOrderTraverse(bst.root, f)
}

func preOrderTraverse(node *Node, f func(interface{})) {
	if node != nil {
		f(node.value)
		preOrderTraverse(node.left, f)
		preOrderTraverse(node.right, f)
	}
}

// PostOrderTraverse 遍历 -- 左右中
func (bst *BinarySearchTree) PostOrderTraverse(f func(interface{})) {
	postOrderTraverse(bst.root, f)
}

func postOrderTraverse(node *Node, f func(interface{})) {
	if node != nil {
		postOrderTraverse(node.left, f)
		postOrderTraverse(node.right, f)
		f(node.value)
	}
}

// Min 最小值
func (bst *BinarySearchTree) Min() interface{} {
	node := bst.root
	if node == nil {
		return nil
	}
	for {
		if node.left == nil {
			return node.value
		}
		node = node.left
	}
}

// Max 最大值
func (bst *BinarySearchTree) Max() interface{} {
	node := bst.root
	if node == nil {
		return nil
	}
	for {
		if node.right == nil {
			return node.value
		}
		node = node.right
	}
}

// Search 搜索
func (bst *BinarySearchTree) Search(key string) bool {
	return search(bst.root, key)
}

func search(node *Node, key string) bool {
	if node == nil {
		return false
	}
	if node.key < key {
		return search(node.right, key)
	}
	if node.key > key {
		return search(node.left, key)
	}
	return true
}

// Remove 删除
func (bst *BinarySearchTree) Remove(key string) {
	remove(bst.root, key)
}

func remove(node *Node, key string) *Node {
	if node == nil {
		return nil
	}
	if node.key > key {
		node.left = remove(node.left, key)
		return node
	}
	if node.key < key {
		node.right = remove(node.right, key)
		return node
	}
	// key == node.key
	if node.left == nil && node.right == nil { // 节点没有子节点，直接置nil，即可删除
		node = nil
		return nil
	}
	if node.left == nil { // 左节点为nil，用右节点替换该节点，即可删除
		node = node.right
		return node
	}
	if node.right == nil { // 右节点为nil，用左节点替换该节点，即可删除
		node = node.left
		return node
	}
	// 左右节点均存在
	leftMostRightSide := node.right // 从右节点开始
	for {
		// 找出右边最小的值，该值将被移动到被删除节点的位置
		if leftMostRightSide != nil && leftMostRightSide.left != nil {
			leftMostRightSide = leftMostRightSide.left
		} else {
			break
		}
	}
	node.key, node.value = leftMostRightSide.key, leftMostRightSide.value // 替换
	node.right = remove(node.right, node.key)                             // 左节点不变，右节点更新，因为从右子树拿了一个最小值过来填充这个位置
	return node
}

// String 输出字符串 -- 从左往右，从上往下
func (bst *BinarySearchTree) String() {
	stringify(bst.root, 0)
}

func stringify(node *Node, level int) {
	if node != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "    "
		}
		format += "---[ "
		level++
		stringify(node.left, level)
		fmt.Printf(format+"%v\n", node.key)
		stringify(node.right, level)
	}
}

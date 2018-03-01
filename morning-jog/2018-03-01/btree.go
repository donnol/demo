package p20180301

// b-tree 维基 https://zh.wikipedia.org/wiki/B%E6%A0%91
// 代码来源 github.com/emirpasic/gods

// 假设有以下树
//  1 3 5 7 9 -- 在同一个节点内
// 0 2 4 6 8 10 -- 每个键在不同节点，父节点都是上面的节点

// 在一个节点里搜索指定key
// 成功找到时，返回key的下标；搜索失败的话，返回小于key的最大值的下标
func (tree *Tree) search(node *Node, key interface{}) (index int, found bool) {
	// 假设节点里面含有5个元素，键分别是 1 3 5 7 9
	// 待搜索元素键是 6

	low, high := 0, len(node.Entries)-1 // low=0 high=4
	var mid int
	for low <= high {
		mid = (high + low) / 2                                 // 取中间位置下标 mid=2 键是 5
		compare := tree.Comparator(key, node.Entries[mid].Key) // 比较待搜索键和中间位置键，显然 6 > 5
		switch {
		case compare > 0:
			low = mid + 1 // low=3
		case compare < 0:
			high = mid - 1
		case compare == 0:
			return mid, true
		}
	}
	return low, false // low=3
}

// searchRecursively searches recursively down the tree starting at the startNode
func (tree *Tree) searchRecursively(startNode *Node, key interface{}) (node *Node, index int, found bool) {
	if tree.Empty() {
		return nil, -1, false
	}
	node = startNode
	for {
		index, found = tree.search(node, key)
		if found {
			return node, index, true
		}
		if tree.isLeaf(node) {
			return nil, -1, false
		}
		node = node.Children[index]
	}
}

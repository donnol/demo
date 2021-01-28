package merkletree

import "crypto/sha256"

// 由一个根节点（root）、一组中间节点和一组叶节点（leaf）组成。
// 叶节点（leaf）包含存储数据或其哈希值，中间节点是它的两个孩子节点内容的哈希值，
// 根节点也是由它的两个子节点内容的哈希值组成。所以Merkle树也称哈希树。
//
// 特点：
// 	叶节点存储的是数据文件，而非叶节点存储的是其子节点的哈希值（Hash，通过SHA1、SHA256等哈希算法计算而来），这些非叶子节点的Hash被称作路径哈希值（可以据其确定某个叶节点到根节点的路径）,
// 	叶节点的Hash值是真实数据的Hash值。因为使用了树形结构, 其查询的时间复杂度为 O(logn)，n是节点数量。
//
// 默克尔树的另一个特点是，底层数据的任何变动，都会传递到其父节点，一直到树根。
//
// 映像中go module也是用的MerkleTree
type MerkleTree struct {
	root *node
}

type node struct {
	hvalue []byte // 哈希值
	data   []byte // 数据，只有叶子节点有数据
	childs []node // 子节点
}

func get() *node {

	return nil
}

func add() {

}

func mod() {

}

func del() {

}

var (
	h = sha256.New()
)

func hash(data []byte) []byte {
	_, err := h.Write(data)
	if err != nil {
		panic(err)
	}
	sum := h.Sum(nil)

	h.Reset()

	return sum
}

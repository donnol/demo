package p20171228

// Tree 树
type Tree struct {
	root    *node
	numnode int // 节点数量
}

// Get 获取键的值
func (t *Tree) Get(key interface{}) (value interface{}, found bool) {
	return
}

// Add 添加键值
func (t *Tree) Add(key, value interface{}) (success bool) {
	return
}

// Mod 修改键值
func (t *Tree) Mod(key, value interface{}) (success bool) {
	return
}

// Del 删除键值
func (t *Tree) Del(key interface{}) (success bool) {
	return
}

// Print 打印
func (t *Tree) Print() {

}

// NodeNum 节点数
func (t *Tree) NodeNum() (num int) {
	return
}

// node 节点
type node struct {
	parent *node
	key    interface{}
	value  interface{}
	child  []*node
}

// 在这个节点里搜索 key
func (n *node) search(key interface{}) (index int, found bool) {
	// 本节点
	if n.key == key {
		index = -1
		found = true
		return
	}

	// 遍历子节点
	low, high := 0, len(n.child)-1
	for low <= high {
		mid := (low + high) / 2
		tmp := compare(n.child[mid].key, key)
		switch tmp {
		case 1:
			high = mid
		case -1:
			low = mid
		case 0:
			index = mid
			found = true
			return
		}
	}
	index = low

	return
}

func (n *node) add(key, value interface{}) (success bool) {
	return
}

func (n *node) del(key interface{}) (success bool) {
	return
}

func (n *node) mod(key, value interface{}) (success bool) {
	return
}

func compare(l, r interface{}) int {
	if l.(string) > r.(string) {
		return 1
	} else if l.(string) < r.(string) {
		return -1
	} else {
		return 0
	}
}

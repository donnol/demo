package uniontree

/*
并查集，判连通用
*/
type UnionFindSet struct {
	father  []int // 存储结点的父亲
	nodeNum int   // 总结点个数
}

func NewUnionSet(n int) *UnionFindSet {
	us := &UnionFindSet{}
	us.nodeNum = n + 1
	us.father = make([]int, us.nodeNum)
	for i := range us.father {
		us.father[i] = i
	}
	return us
}

// 查询父结点
func (us *UnionFindSet) Find(x int) int {
	for us.father[x] != x {
		x = us.father[x]
	}
	return x
}

//合并结点
func (us *UnionFindSet) Union(x, y int) bool {
	x = us.Find(x)
	y = us.Find(y)
	if x == y {
		return false
	}
	us.father[x] = y
	return true
}

func FindCircleNum(M [][]int) int {
	us := NewUnionSet(len(M))
	//fmt.Printf("input: %+v\n", M)
	//fmt.Printf("union set: %+v\n", us)
	//fmt.Println("===")

	for i, row := range M {
		for j, isFriend := range row {
			if isFriend == 1 {
				us.Union(i, j)
			} // 是朋友则合并
		}
	}

	friendCircle := 0
	for i := 0; i < len(M); i++ {
		if us.Find(i) == i {
			friendCircle++
		} // 查看集合代表个数
	}

	return friendCircle
}

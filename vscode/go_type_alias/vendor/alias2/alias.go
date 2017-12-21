package alias2

// M M 原始类型
type M string

func (m M) String() string {
	return string(m)
}

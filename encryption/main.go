package main

func main() {

}

// Moss 摩斯密码
type Moss struct {
	Dit string `json:"·"` // 点
	Dah string `json:"-"` // 划
}

// String 文本输出
func (m *Moss) String() string {
	return ""
}

// Translate 翻译
func (m *Moss) Translate(typ string, text string) (string, error) {
	return "", nil
}

// Mosser 摩斯密码接口
type Mosser interface {
	Translate(typ string, text string) (string, error)
}

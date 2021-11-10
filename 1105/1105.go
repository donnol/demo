package main

type Model struct {
}

var (
	conn interface{}
)

var (
	batchNum int
)

func batchInsert(data []model.Model) error {
	// 校验

	// 看data数组的长度，
	// 分批插入
	l := len(data) // 201

	var start int // 0
	for {
		if start > l {
			break
		}
		end := start + batchNum // 100 200 300
		left := l - start       // 201 101 1
		if left < batchNum {
			end = start + left // 201
		}

		tmp := data[start:end] // 0 100 | 100 200 | 200 201
		_ = tmp

		for i := 0; i < 5; i++ {
			// 外部存储
			// conn tmp

			if err != nil {
				// log

				n := 1 + pow(2, i)
				// sleep
				time.Sleep(n) // 1秒，2^i

				continue
			}

			break
		}

		start += batchNum // 100 | 200
	}

	return nil
}

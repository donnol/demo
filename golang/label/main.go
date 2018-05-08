package main

import (
	"fmt"
)

func main() {
	fmt.Println("break label test ======")
	breakLabel()

	fmt.Println("goto label test ======")
	gotoLabel()

	fmt.Println("continue label test ======")
	continueLabel()
}

func breakLabel() {
label:
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 3 {
			break label // 语句执行后，不再打印后面的 4，可将语句注释来作对比
		}
	}
	// label 声明在 break 后，报错：break label not defined: label
	// 	label label defined and not used
	// label:
	fmt.Println("finish loop")
}

func gotoLabel() {
	// label1:
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 3 {
			// goto label1 // 会出现死循环，一直打印 0 1 2 3
			goto label2 // 跳过 4，效果与 break label1 类似，可用 break label1 替代
			// break label1
		}
	}
label2:
	fmt.Println("finish loop")
}

func continueLabel() {
label:
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 3 {
			continue label // 跳过该次循环后面的语句，开始下一次循环
		}
		fmt.Println("second:", i)
	}
	// label 声明在 break 后，报错：break label not defined: label
	// 	label label defined and not used
	// label:
	fmt.Println("finish loop")
}

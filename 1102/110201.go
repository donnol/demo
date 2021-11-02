// 股票每天的价格是arr（这是一个列表），你会选择1天买入股票，选之后的另外1天卖出，所以需要一个函数在这种操作下可以把利润最大化
// 如：arr = [7,1,5,3,6,4]
package main

func main() {
	r := test2([]int{7, 1, 5, 3, 6, 4})
	println(r)
}

func test(arr []int) int {
	var r int

	l := len(arr)
	for i := 0; i < l-1; i++ {
		buy := arr[i]
		for j := i + 1; j < l; j++ {
			sold := arr[j]

			diff := sold - buy
			if diff > r {
				r = diff
			}
		}
	}

	return r
}

func test2(arr []int) int {
	var r int

	l := len(arr)

	// 遍历的同时，记录前面部分里最小的值
	// 因为第i个卖出所能得到的最大收益就是前面里最小的买入
	var min int
	for i := 0; i < l; i++ {

		if i != 0 {
			diff := arr[i] - min
			if diff > r {
				r = diff
			}
		}

		if i == 0 || arr[i] < min {
			min = arr[i]
		}

	}

	return r
}

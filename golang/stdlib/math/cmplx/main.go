package main

import (
	"log"
	"math"
	"math/cmplx"
)

func main() {
	for _, x := range []complex128{
		0,
		0i,
		1,
		-1,
		1i,
		-1i,
		1 + 1i,
		-1 + 1i,
		3 + 4i,
		-3 + 4i,
		3 - 4i,
		-3 - 4i,
		1i * math.Pi,
	} {
		log.Printf("===== cmplx is %v =====\n", x) // 分隔符
		log.Println("Abs: ", cmplx.Abs(x))         // 求复数的绝对值，即三角形的斜边长
		log.Println("Cos: ", cmplx.Cos(x))         // 求复数的余弦值
		log.Println("Sin: ", cmplx.Sin(x))         // 求复数的正弦值
		log.Println("Exp: ", cmplx.Exp(x))         // e**x，e 的 x 次幂
		log.Println("Log: ", cmplx.Log(x))         // x 的自然对数
		log.Println("Phase: ", cmplx.Phase(x))     // 相位(幅角)，取值范围 [-π, π]
		r, θ := cmplx.Polar(x)                     // x = r * e**θi，r 是绝对值，θ 是相位
		log.Println("Polar: ", r, θ)
		log.Println("Rect: ", cmplx.Rect(r, θ)) // 使用极坐标计算
	}

	// 欧拉公式的特殊形式为 e**(i*pi) + 1 = 0，e 的 iπ 次幂 加 1 等于 0
	log.Printf("Euler's identity: %.1f\n", cmplx.Exp(1i*math.Pi)+1)

	r, theta := cmplx.Polar(2i)
	log.Printf("r: %.1f, θ: %.1f*π", r, theta/math.Pi)
}

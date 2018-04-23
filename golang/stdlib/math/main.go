package main

import (
	"log"
	"math"
)

func main() {
	x, y := 1.0, 2.0
	min := math.Min(x, y)
	log.Println(min) // 1

	log.Println(math.Cos(0)) // 1

	log.Println(math.Inf(1))  // +Inf
	log.Println(math.Inf(0))  // +Inf
	log.Println(math.Inf(-1)) // -Inf

	log.Println(math.NaN())    // NaN
	log.Println(math.Acos(2))  // NaN
	log.Println(math.Acos(-2)) // NaN

	// x的立方根
	log.Println(math.Cbrt(x)) // 1
	log.Println(math.Cbrt(y)) // 1.2599210498948732

	// 大于或等于x的最小整数
	log.Println(math.Ceil(-0.9)) // -0
	log.Println(math.Ceil(0.9))  // 1
	log.Println(math.Ceil(1.9))  // 2
	log.Println(math.Ceil(2.0))  // 2

	// x的值，y的符号
	log.Println(math.Copysign(1.1, 2.2))   // 1.1
	log.Println(math.Copysign(1.1, -2.2))  // -1.1
	log.Println(math.Copysign(-1.1, 2.2))  // 1.1
	log.Println(math.Copysign(-1.1, -2.2)) // -1.1

	// x-y的最大值或0(结果是大于等于0的数)
	log.Println(math.Dim(2.2, 1.1))   // 1.1
	log.Println(math.Dim(-2.2, 1.1))  // 0
	log.Println(math.Dim(2.2, -1.1))  // 3.3000000000000003
	log.Println(math.Dim(-2.2, -1.1)) // 0
	log.Println(math.Dim(1.1, 2.2))   // 0

	// 返回x的错误函数？
	log.Println(math.Erf(2.2))     // 0.9981371537020182
	log.Println(math.Erfc(2.2))    // 0.0018628462979818898
	log.Println(math.Erfcinv(2.2)) // NaN
	log.Println(math.Erfinv(2.2))  // NaN

	// e的x指数幂: e**x
	log.Println(math.Exp(2.2)) // 9.025013499434122
	// 2的x指数幂：2**x
	log.Println(math.Exp2(2.2)) // 4.59479341998814
	// e的x-1指数幂: e**x - 1
	log.Println(math.Expm1(2.2)) // 8.025013499434122

	// x的二进制表示
	log.Println(math.Float32bits(2.2))                     // 1074580685
	log.Println(math.Float32frombits(1074580685))          // 2.2
	log.Println(math.Float64bits(2.2))                     // 4612136378390124954
	log.Println(math.Float64frombits(4612136378390124954)) // 2.2

	// 小于或等于x的最大整数值
	log.Println(math.Floor(2.2))  // 2
	log.Println(math.Floor(-2.2)) // -3

	// 将x分解为frac乘以2的exp指数幂: x == frac X 2**exp
	log.Println(math.Frexp(2.2)) // frac = 0.55 exp = 2

	// x的伽马函数(https://zh.wikipedia.org/wiki/%CE%93%E5%87%BD%E6%95%B0)
	log.Println(math.Gamma(2.2)) // 1.101802490879713

	// 三角形的斜边长: x^2+y^2的开方
	log.Println(math.Hypot(1.0, 2.0)) // 2.23606797749979
	log.Println(math.Hypot(3.0, 4.0)) // 5

	// x的二进制指数
	log.Println(math.Ilogb(3.3)) // 1
	log.Println(math.Ilogb(2.2)) // 1
	log.Println(math.Ilogb(1.0)) // 0

	// 贝塞尔函数
	log.Println(math.J0(2.2))    // 0.11036226692217384
	log.Println(math.J1(2.2))    // 0.555963049819064
	log.Println(math.Jn(2, 2.2)) // 0.3950586874587934

	// Frexp的逆运算
	log.Println(math.Ldexp(0.55, 2)) // 2.2

	// x的伽马函数的自然对数和符号
	log.Println(math.Lgamma(2.2)) // 0.09694746679063887 1

	// x的自然对数，以e为底
	log.Println(math.Log(2.2))             // 0.7884573603642703
	log.Println(math.Log(math.Gamma(2.2))) // 0.09694746679063906

	// x的对数，以10为底
	log.Println(math.Log10(2.2)) // 0.3424226808222063
	log.Println(math.Log10E)     // 0.4342944819032518

	// x+1的自然对数
	log.Println(math.Log1p(2.2)) // 1.1631508098056809

	// x的二进制对数，以2为底
	log.Println(math.Log2(2.2)) // 1.1375035237499351

	// x的二进制指数
	log.Println(math.Logb(2.2)) // 1
}

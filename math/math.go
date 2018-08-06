package math

import "math"

//float64 取小数点后面n位

//计算结果取小数的n位
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

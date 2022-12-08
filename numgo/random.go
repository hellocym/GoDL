package numgo

import "math/rand"

func Randn(dims ...int) NDArray {
	size := 1
	for _, dim := range dims {
		size *= dim
	}
	temp := NDArray{dims, make([]float64, size)}
	for i := 0; i < size; i++ {
		temp.Data[i] = rand.NormFloat64()
	}
	return temp
}

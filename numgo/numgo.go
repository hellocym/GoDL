package numgo

import "fmt"

type NDArray struct {
	Shape []int
	Data  []float64
}

// Zeros Create a new NDArray filled with zeros.
func Zeros(dims ...int) NDArray {
	size := 1
	for _, dim := range dims {
		size *= dim
	}
	return NDArray{dims, make([]float64, size)}
}

// Builtin function String() for NDArray, returns formatted string of a NDArray.
func (a NDArray) String() string {
	res := ""
	if len(a.Shape) == 1 {
		res += "["
		sep := ""
		for _, val := range a.Data {
			res += sep
			res += fmt.Sprintf("%f", val)
			sep = ", "
		}
		res += "]"
	} else {
		res += "["
		sep := ""
		for i := 0; i < a.Shape[0]; i++ {
			size := 1
			for _, dim := range a.Shape[1:] {
				size *= dim
			}
			temp := NDArray{a.Shape[1:], a.Data[i*size : (i+1)*size]}
			res += sep
			res += temp.String()
			sep = ", "
		}
		res += "]"
	}
	return res
}

// Iloc get an element of a NDArray with indices
// a.Shape = (3, 4, 5), a[1, 2, 3] = a.Data[1*4*5 + 2*5 + 3]
// a.Shape = (3, 4, 5, 6) a[1, 2, 3, 4] = a.Data[1*4*5*6 + 2*5*6 + 3*6 + 4]
func (a NDArray) Iloc(indices ...int) float64 {
	index := 0
	for i := 0; i < len(indices); i++ {
		temp := 1
		for j := i + 1; j < len(a.Shape); j++ {
			temp *= a.Shape[j]
		}
		index += indices[i] * temp
	}
	return a.Data[index]
}

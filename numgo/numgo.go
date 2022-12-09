package numgo

import (
	"fmt"
)

type NDArray struct {
	Shape []int
	Data  []float64
}

// Size Return the number of elements in the NDArray.
func (a NDArray) Size() int {
	size := 1
	for _, dim := range a.Shape {
		size *= dim
	}
	return size
}

// Sum compute the sum of NDArray elements.
func Sum(a NDArray, axis int) NDArray {
	if axis < 0 || axis >= len(a.Shape) {
		panic("axis out of range")
	}
	dimAccum := a.Shape[axis]
	resSize := a.Size() / dimAccum
	resData := make([]float64, resSize)
	resShape := make([]int, len(a.Shape))
	copy(resShape, a.Shape)
	resShape[axis] = 1

	if axis == 0 {
		for i := 0; i < resSize; i++ {
			for j := 0; j < dimAccum; j++ {
				resData[i] += a.Data[i+j*resSize]
			}
		}
	} else {
		for i := 0; i < resSize; i++ {
			resIndices := index2indices(i, resShape)
			for j := 0; j < dimAccum; j++ {
				resIndices[axis] = j
				resData[i] += a.Data[indices2index(resIndices, a.Shape)]
			}
		}
	}
	resShape = append(resShape[:axis], resShape[axis+1:]...)
	res := NDArray{resShape, resData}
	return res
}

// Prod compute the product of NDArray elements.
func Prod(a NDArray, axis int) NDArray {
	if axis < 0 || axis >= len(a.Shape) {
		panic("axis out of range")
	}
	dimAccum := a.Shape[axis]
	resSize := a.Size() / dimAccum
	resData := make([]float64, resSize)
	resShape := make([]int, len(a.Shape))
	copy(resShape, a.Shape)
	resShape[axis] = 1

	if axis == 0 {
		for i := 0; i < resSize; i++ {
			resData[i] = 1
			for j := 0; j < dimAccum; j++ {
				resData[i] *= a.Data[i+j*resSize]
			}
		}
	} else {
		for i := 0; i < resSize; i++ {
			resIndices := index2indices(i, resShape)
			resData[i] = 1
			for j := 0; j < dimAccum; j++ {
				resIndices[axis] = j
				resData[i] *= a.Data[indices2index(resIndices, a.Shape)]
			}
		}
	}
	resShape = append(resShape[:axis], resShape[axis+1:]...)
	res := NDArray{resShape, resData}
	return res
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

func indices2index(indices []int, shape []int) int {
	index := 0
	for i := 0; i < len(indices); i++ {
		temp := 1
		for j := i + 1; j < len(shape); j++ {
			temp *= shape[j]
		}
		index += indices[i] * temp
	}
	return index
}

func index2indices(index int, shape []int) []int {
	indices := make([]int, len(shape))
	for i := 0; i < len(shape); i++ {
		temp := 1
		for j := i + 1; j < len(shape); j++ {
			temp *= shape[j]
		}
		indices[i] = index / temp
		index %= temp
	}
	return indices
}

// Repeat a NDArray along an axis
// a.Shape = (3, 4, 5), a.Repeat(2, axis=0).Shape = (6, 4, 5)
// a.Shape = (3, 4, 5), a.Repeat(2, axis=1).Shape = (3, 8, 5)
// a.Shape = (3, 4, 5), a.Repeat(2, axis=2).Shape = (3, 4, 10)
func Repeat(a NDArray, repeats int, axis int) NDArray {
	if axis < 0 || axis >= len(a.Shape) {
		panic("axis out of range")
	}
	newShape := make([]int, len(a.Shape))
	copy(newShape, a.Shape)
	newShape[axis] *= repeats
	newData := make([]float64, a.Size()*repeats)
	if axis == 0 {
		for i := 0; i < repeats; i++ {
			copy(newData[i*a.Size():], a.Data)
		}
	} else {
		repeatSize := 1
		for _, dim := range a.Shape[axis:] {
			repeatSize *= dim
		}
		for i := 0; i < a.Size()/repeatSize; i++ {
			temp := a.Data[i*repeatSize : (i+1)*repeatSize]
			for j := 0; j < repeats; j++ {
				copy(newData[(i*repeats+j)*repeatSize:(i*repeats+j+1)*repeatSize], temp)
			}
		}
	}
	res := NDArray{newShape, newData}
	return res
}

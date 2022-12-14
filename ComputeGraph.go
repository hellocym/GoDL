package main

import (
	"GoDL/numgo"
	"fmt"
)

// z = x + y -> dz/dx = 1, dz/dy = 1
// z = x - y -> dz/dx = 1, dz/dy = -1
// z = x * y -> dz/dx = y, dz/dy = x
// z = x / y -> dz/dx = 1/y, dz/dy = -x/y^2
// x1 = x0; x2 = x0 -> dL/dx0 = dL/dx1 + dL/dx2

func main() {
	//D, N := 8, 7
	//x := numgo.Randn(1, D)
	//fmt.Println(x) //print(x)
	//y := numgo.Repeat(x, N, axis=0)
	x := numgo.Randn(2, 3, 4)
	//fmt.Println(x)
	y := numgo.Repeat(x, 2, 2)
	fmt.Println(y)
}

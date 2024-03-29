package println

import (
	"fmt"
)

func main() {
	var x []int
	var y []int
	var z []int

	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	y = append(x, 4)
	z = append(x, 5)

	fmt.Println(y, z)
}

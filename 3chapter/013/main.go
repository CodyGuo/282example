/*
 * @example: 任意次方后的最后三位
 * @author: cody.guo
 */

package main

import (
	"fmt"
)

func main() {
	var x, y, z int = 1, 1, 1
	fmt.Println("请输入两个数, x 和 y:")
	fmt.Scanf("%d %d", &x, &y)
	for i := 1; i <= y; i++ {
		z = z * x % 1000
	}
	if z >= 100 {
		fmt.Printf("%d^%d的最后三位是: %d\n", x, y, z)
	} else {
		fmt.Printf("%d^%d的最后三位是: 0\n", x, y)
	}

}

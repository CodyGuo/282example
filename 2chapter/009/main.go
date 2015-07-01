/*
 * @example: 求10!
 * @author: cody.guo
 */

package main

import (
	"fmt"
)

func main() {

	sum := 0
	for i := 0; i <= 10; i++ {
		if i == 0 || i == 1 {
			sum = 1 // i等于0和1的时候10的阶乘是1
		} else {
			sum *= i
		}
	}
	fmt.Printf("factorial of 10 is: %.2f\n", float32(sum))
}

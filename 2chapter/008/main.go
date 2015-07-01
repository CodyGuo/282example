/*
 * @example: 一个简单的求和程序
 * @author: cody.guo
 */

package main

import (
	"fmt"
)

func main() {
	sum := 0
	a, b := 10, 20
	sum = a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)

}

/*
 * @example: 一个完成的Go语言程序
 * @author: cody.guo
 */
package main

import (
	"fmt"
)

func main() {
	var a, b, sum int = 10, 20, 0
	sum = a + b
	fmt.Printf("sum = a + b, %d + %d = %d\n", a, b, sum)
}

/*
 * @example: 3个数由小到大排序
 * @author: cody.guo
 */

package main

import (
	"fmt"
)

func sort(a, b, c int) {
	var tmp int
	if a > b {
		tmp = a
		a = b
		b = tmp

	}
	if a > c {
		tmp = a
		a = c
		c = tmp

	}
	if b > c {
		tmp = b
		b = c
		c = tmp
	}
	fmt.Println(a, b, c)

}

func main() {
	sort(10, 2, 20)
	sort(2, 300, 100)
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	sort(a, b, c)

}

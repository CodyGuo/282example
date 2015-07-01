/*
 * @example: 阳阳买苹果
 * @author: cody.guo
 */

package main

import (
	"fmt"
)

const (
	price = 0.8
)

func main() {
	// var sum float32
	var day, num int = 1, 2
	var money, ave float32 = price * float32(num), 0
	for num <= 100 {
		if (num * 2) < 100 {
			money += price * float32(num)
			day++
			num *= 2

		} else {
			break
		}

	}
	fmt.Println(day, num)
	ave = money / float32(day)
	fmt.Println(ave)

}

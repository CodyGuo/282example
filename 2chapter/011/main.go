/*
 * @example: 猴子吃桃
 * @author: cody.guo
 */
package main

import (
	"fmt"
)

func main() {
	var day, sum1, sum2 int
	day = 9  // 从第九天开始计算,每天吃前一天的两倍
	sum1 = 1 // 第10天剩下的一个
	for day > 0 {
		sum2 = (sum1 + 1) * 2
		sum1 = sum2
		day--
	}
	fmt.Println(sum1)
}

/*
 * @example: 计算某日是该年的第几天
 * @author: cody.guo
 */
package main

import (
	"fmt"
)

func leap(a int) (result int) {
	// 计算输入的年份是瑞年还是平年
	if a%4 == 0 && a%100 != 0 || a%400 == 0 {
		return 1
	} else {
		return 0
	}
}

func number(year, m, d int) (sum int) {
	a := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31} // 存放平年每月的天数
	b := [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31} // 存放闰年每月的天数

	if leap(year) == 1 {
		for i := 0; i < (m - 1); i++ {
			sum += b[i]
		}
	} else {
		for i := 0; i < (m - 1); i++ {
			sum += a[i]
		}
	}
	sum += d
	return sum

}
func main() {
	var y, m, d int
	fmt.Println("请输入要计算天数的日期,2015 7 1")
	fmt.Scanf("%d %d %d\n", &y, &m, &d)
	fmt.Printf("%d年%d月%d日是今年的第%d天.\n", y, m, d, number(y, m, d))

}

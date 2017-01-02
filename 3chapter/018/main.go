/**
 * @example: 判断三角形的类型
 * @author: cody.guo
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	var s, area float64

	fmt.Print("请输入三条边:")
	fmt.Scanf("%f %f %f", &a, &b, &c)
	if a+b > c && b+c > a && a+c > b {
		s = (a + b + c) / 2
		area = math.Sqrt(float64(s * (s - a) * (s - b) * (s - c)))
		fmt.Printf("面积是: %f\n", area)

		if a == b && a == c {
			fmt.Println("等边三角形")
		} else if a == b || a == c || b == c {
			fmt.Println("等腰三角形")
		} else if (a*a+b*b == c*c) || (a*a+c*c == b*b) ||
			(b*b+c*c == a*a) {
			fmt.Println("直角三角形")
		} else {
			fmt.Println("普通三角形")
		}
	} else {
		fmt.Println("不成形")
	}

}

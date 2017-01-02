/**
 * 希尔排序
 */
package main

import "fmt"

func shshort(s [10]int) {
	n := len(s)
	d := n / 2
	for d >= 1 {
		for i := d + 1; i >= 0 && i < n; i++ {
			lookout := s[i]
			j := i - d
			for ; j >= 0 && lookout < s[j]; j -= d {
				s[j+d] = s[j]
			}
			s[j+d] = lookout
		}
		d = d / 2
	}
	fmt.Printf("插入数据排序后顺序：\n%v\n", s)
}

func main() {
	var a [10]int
	fmt.Println("请输入10个数据: ", len(a))
	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Printf("原始顺序：\n%v\n", a)
	shshort(a)
}

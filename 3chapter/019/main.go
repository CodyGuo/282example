/**
 * 直接插入排序
 */
package main

import "fmt"

func insort(s [10]int) {
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			// 给监视哨赋值
			lookout := s[i]
			j := i - 1
			for ; j >= 0 && s[j] > lookout; j-- {
				// 数据右移
				s[j+1] = s[j]
			}
			// 在确定位置插入监视哨
			s[j+1] = lookout
		}
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
	insort(a)
}

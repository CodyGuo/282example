/**
 * 冒泡排序
 */
package main

import "fmt"

func msort(nums [10]int) [10]int {
	for i, n := 0, len(nums)-1; i >= 0 && i < n; i++ {
		for j := 0; j >= 0 && j < n-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				fmt.Printf("swap -> %v\n", nums)
			}
		}
	}
	return nums
}

func main() {
	nums := inputNum()
	fmt.Printf("原始顺序: \n  %v\n", nums)
	nums = msort(nums)
	fmt.Printf("排序后顺序: \n  %v\n", nums)
}

func inputNum() [10]int {
	var nums [10]int
	fmt.Println("请输入10个数据: ")
	for i, n := 0, len(nums); i < n; i++ {
		fmt.Scan(&nums[i])
	}
	return nums
}

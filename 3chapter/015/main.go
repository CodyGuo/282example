/*
 * @example: 婚礼上的谎言
 * @author: cody.guo
 */
package main

import (
	"fmt"
)

func main() {
	for A := 1; A <= 3; A++ { // 穷举A的所有可能
		for B := 1; B <= 3; B++ { // 穷举B的所有可能
			for C := 1; C <= 3; C++ { // 穷举C的所有可能
				if A != 1 && C != 1 && C != 3 && A != B && A != C && B != C {
					fmt.Printf("%c 将嫁给 A\n", 'X'+A-1)
					fmt.Printf("%c 将嫁给 B\n", 'X'+B-1)
					fmt.Printf("%c 将嫁给 C\n", 'X'+C-1)
				}
			}
		}
	}

}

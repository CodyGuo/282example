/*
 * @example: 百元买百鸡
 * @author: cody.guo
 */
package main

import (
    "fmt"
)

func main() {
    // var cock, hen, chick int

    for cock := 0; cock <= 20; cock++ {
        for hen := 0; hen <= 33; hen++ {
            for chick := 3; chick <= 99; chick++ {
                if (5*cock + 3*hen + chick/3) == 100 {
                    if (cock + hen + chick) == 100 {
                        if (chick % 3) == 0 {
                            fmt.Println("公鸡:", cock, "母鸡:", hen, "小鸡:", chick)
                        }
                    }
                }
            }
        }
    }
}

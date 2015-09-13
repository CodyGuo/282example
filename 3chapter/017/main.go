/*
 * @example: 打渔晒网问题
 * @author: cody.guo
 */
package main

import (
    "fmt"
    // "time"
)

type Day struct {
    aYear []int
    bYear []int
}

//{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
type Year struct {
    days Day
}

func (yy Year) sumYear(y int) (day int) {
    if (y%4 == 0) && (y%100 != 0) || (y%400 == 0) {
        for i := 0; i < len(yy.days.aYear); i++ {
            day = day + yy.days.aYear[i]
        }
        fmt.Println(y, "sumyear 润年")
        return day
    }
    for i := 0; i < len(yy.days.bYear); i++ {
        day = day + yy.days.bYear[i]
    }
    fmt.Println(y, "sumyear 平年")
    return day
}

func (mm Year) sumOther(y, m, d int) (day int) {
    if (y%4 == 0) && (y%100 != 0) || (y%400 == 0) {
        for i := 0; i < m-1; i++ {
            day = day + mm.days.aYear[i]
        }
        fmt.Println(y, "other 润年")
        return day + d
    }
    for i := 0; i < m-1; i++ {
        day = day + mm.days.bYear[i]
    }
    fmt.Println(y, "other 平年")
    return day + d
}

func main() {
    years := new(Year)
    years.days.aYear = []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    years.days.bYear = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    var sum int
    y := 2011
    m := 1
    d := 6
    for i := 2011; i < y; i++ {
        sum += years.sumYear(i)
    }

    sum += years.sumOther(y, m, d)

    fmt.Println(sum)
    fmt.Println(years.days.aYear, years.days.bYear)

    if sum%5 < 4 && sum%5 > 0 {
        fmt.Println(y, m, d, "打渔")
    } else {
        fmt.Println(y, m, d, "晒网")
    }

}

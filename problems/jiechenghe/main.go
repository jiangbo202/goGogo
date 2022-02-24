/**
 * @Author: jiangbo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/02/24 9:30 下午
 */

package main

import "fmt"

func sum(x int) uint64 {
    var s uint64 = 1
    for i := 0; i <= x; i++ {
        s1 := s * uint64(i)
        s += s1
    }
    return s
}

func main() {
    var m int
    fmt.Scanf("%d", &m)

    fmt.Println(sum(m))
}

// 计算n的阶乘只之和
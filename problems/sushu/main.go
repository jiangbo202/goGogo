/**
 * @Author: jiangbo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/02/24 9:14 下午
 */

package main

import "fmt"

/*
    输出100-200之间的素数
 */

func isPrime(x int) bool {
    for i := 2; i < x; i++ {
        if x%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    var m, n int
    fmt.Scanf("%d%d", &m, &n)

    for i := m; i < n; i++ {
        if isPrime(i){
            fmt.Printf("%d is Prime.\n", i)
        }
    }
}

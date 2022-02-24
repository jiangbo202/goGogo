/**
 * @Author: jiangbo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/02/24 9:21 下午
 */

package main

import "fmt"

func isTarget(x int) bool {
    var i, j, k int
    i = x % 10
    j = (x / 10) % 10
    k = (x / 100) % 10   // 这里也可以转成字符串，然后一个一个摘下来(按照索引)

    sum := i*i*i + j*j*j + k*k*k
    return sum == x
}

func main() {
    var m, n int
    fmt.Scanf("%d%d", &m, &n)

    for i := m; i < n; i++ {
        if isTarget(i) {
            fmt.Printf("%d is shuixianhua\n", i)
        }
    }
}

// 得到100-999之间的水仙花数: 三位数，每一个位的立方相加等于自身
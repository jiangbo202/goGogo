/**
 * @Author: jiangbo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/02/24 9:53 下午
 */

package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println(time.Now().Format("2006/01/02 15:04:05"))
    fmt.Println(time.Now().UnixNano())
    fmt.Println(time.Now().Unix())
}

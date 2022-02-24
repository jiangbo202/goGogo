/**
 * @Author: jiangbo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/02/24 10:30 下午
 */

package main

import (
    "context"
    "fmt"
    "sync/atomic"
    "time"
)

// https://zhuanlan.zhihu.com/p/343761986

import (
    "golang.org/x/sync/singleflight"
)

type Result string

func find(ctx context.Context, query string) (Result, error) {
    return Result(fmt.Sprintf("result for %q", query)), nil
}

func main() {
    var g singleflight.Group
    const n = 5
    waited := int32(n)
    done := make(chan struct{})
    key := "https://weibo.com/1227368500/H3GIgngon"
    for i := 0; i < n; i++ {
        go func(j int) {
            v, _, shared := g.Do(key, func() (interface{}, error) {
                ret, err := find(context.Background(), key)
                return ret, err
            })
            time.Sleep(3 * time.Second)
            if atomic.AddInt32(&waited, -1) == 0 {
                close(done)
            }

            fmt.Printf("index: %d, val: %v, shared: %v\n", j, v, shared)
        }(i)
    }

    select {
    case <-done:
        fmt.Println("work OK")
    case <-time.After(time.Second):
        fmt.Println("Do hangs")
    }
}

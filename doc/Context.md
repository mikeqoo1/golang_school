# Context 簡介 (用途在於不同的goroutine中傳遞訊息或是取消)

## Context 結構
```
Done方法:在Context被取消或timeout會返回一個close的channel,close的channel可以當廣播通知，告訴给context相關func要停止當前工作return

Err方法:return context為什麼被取消的錯誤原因

Deadline方法:return context何時會timeout

Value方法:return context相關的資料
```

當一個父operation啟動一個goroutine用於子operation，這些子operation不能夠取消父operation.
下面描述的WithCancel func提供一種方式可以取消新建的Context. Context可以安全的被多個goroutine使用。可以把一個Context傳遞给任意多個goroutine然後cancel這個context的时候就能夠通知到所有的goroutine。

範例:
```go
func main() {
    //WithTimeout 等於 WithDeadline
    //ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
    ctx, cancel := context.WithCancel(context.Background())
    go watch(ctx, "【監控1】")
    go watch(ctx, "【監控2】")
    go watch(ctx, "【監控3】")

    //程式的執行時間
    time.Sleep(10 * time.Second)

    fmt.Println("可以了，通知監控停止")
    cancel()

    //為了檢測監控過是否停止，如果没有監控輸出，就表示停止了
    time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
    for {
        select {
            case <-ctx.Done():
                fmt.Println(name, "監控退出，停止了...")
                return
            default:
                fmt.Println(name, "goroutine監控中...")
                time.Sleep(2 * time.Second)
        }
    }
}
```

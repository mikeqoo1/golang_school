# golang_school
golang小教室

# 說明
大概介紹golang的一些處理,或是我有踩到的坑

* Creating Packages
```
要先移倒該lib目錄下執行go install,編譯出類似c的lib連結檔,接下來就可以到main.go中,import後就可以使用


一個目錄下的同級檔案屬於同一個package
package名稱可以與目錄不同名稱, 但盡可能一樣
main package為應用程式執行的entry point; 若沒有main package則無法編譯成可執行的檔案在bin下
package name, Go團隊建議簡單扁平為原則。 所以盡量避免下划線、中划線和參雜大寫字母。
```

# 範例執行
go run main.go就可以
go test -v -bench=. 測試名稱_test.go -benchmem

執行帶參數1跑leetcode 2跑樹狀展示 3跑有的沒的 4跑面試 5跑GTK 6跑排序

# MakeFile說明
make

## GOMOD 說明

1. 切到專案目錄下執行 ```go mod init```

2. 程式寫好後, 直接執行, gomod會自動下載

關於管理的指令

想要先下載lib, 還是可以用以前的go get, 但是在gomod底下, 可以指定版本
像是 ```go get lib@v1.0.5```, 也可以是分支或tag ```go get lib@devel```, 也可以是github的commit id (那些雜湊亂碼)

想下載全部的lib 指令 ```go mod download```

查看所有升级lib版本 指令 ```go list -u -m all```

升級下一個版本或補釘版本 指令 ```go get -u rsc.io/quote```

只升級補丁版本 指令 ```go get -u=patch rscio/quote```

升降版本，可以使用比較運算來處理 指令 ```go get lib@'<v1.6.2'```

程式碼中沒用的到lib, 自動移出gomod 指令 ```go mod tidy```

詳細請參考 - 開始使用 Go Module - isLishude的文章 - 知乎
https://zhuanlan.zhihu.com/p/59687626

## 面試題

- golang的垃圾回收機制 (Garbage Collection)
```
引用計數 (reference counting)

標記-清除法 (mark and sweep)

STW (stop the world)
```
- GC觸發條件
```
1. 超過内存大小szie
2. 達到定時時間
閥值是由一個gcpercent的變量控制的, 當新分配的内存占已在使用中的内存的比例超過gcprecent時就會觸發, 比如一次回收完畢後
内存的使用量為2M, 那下次回收的時機則是内存分配到4M的時候, 也就是說, 並不是内存分配越多, 垃圾回收頻率越高
如果一直達不到内存大小的閥值, 這時候GC就會被定時時間觸發, 比如一直達不到10M, 那就定时（預設2min觸發一次）觸發一次GC保證回收
```
- goroutine 避免 react contain
```
使用sync.Mutex來處理, 但是要注意鎖的解, 和效能問題
使用chennel來溝通
```

- Go深度解析

- https://github.com/qcrao/Go-Questions


## GoTK3的坑

```txt
新版本的glib2（2.68.1.1及以上版本）的接口與gotk3庫的函式不兼容問題
gotk3裡面的
"g_binding_get_source" "g_binding_get_target" "g_binding_get_source_property"  舊版本
分别被
"g_binding_dup_source" "g_binding_dup_target" "g_binding_get_target_property"  新版本./
替換掉了
但是gotk3卻没有修改, 如github.com\gotk3\gotk3@v0.6.1\glib\gbinding.go中还是用了舊的宣告 沒有把廢掉的宣告替换 所以只需要把廢掉的宣告替换就可以正常了
```

## pprof 效能分析

```txt
預先安裝 sudo apt-get install graphviz
分2步
1. 產生數據

2. 分析數據

go tool pprof -http=:8000 profile

會在本地端的8000產生結果

https://juejin.cn/post/7122473470424219656
https://developer.aliyun.com/article/1109652
```

# golang_school
golang小教室

# 說明
大概介紹golang的一些處理,或是我有踩到的坑

* Creating Packages
```
要先移倒該lib目錄下執行go install,編譯出類似c的lib連結檔,接下來就可以到main.go中,import後就可以使用
```

# 範例執行
go run main.go就可以
go test -v -bench=. 測試名稱_test.go -benchmem

#MakeFile說明
make leetcode=1 會編譯出跑LeetCode的程式

make就只跑原本的小功能, 不跑LeetCode

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


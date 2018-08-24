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
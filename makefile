# 執行檔的名稱, 版本, 參數設定
BINARY=GoTeach.out
VERSION=1.0.0
BUILD=`date +%FT%T%z`

# 預設標籤: 編譯程式
build:
	go build -o ${BINARY}

# 安裝專案: copies binaries
install:
	go install

# 清理專案: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
.PHONY:  clean install
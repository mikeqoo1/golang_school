# 執行檔的名稱, 版本, 參數設定
BINARY=GoTeach
VERSION=1.0.0
BUILD=`date +%FT%T%z`

ifeq ($(leetcode),1)
IsLeetCode=TURE
endif

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD} -X main.IsLeetCode=${IsLeetCode}"
# 預設標籤: 編譯程式
build:
	go build ${LDFLAGS} -o ${BINARY}
# 安裝專案: copies binaries
install:
	go install ${LDFLAGS}
# 清理專案: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
.PHONY:  clean install
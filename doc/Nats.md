# Nats Streaming 詳細說明

除了NATS平台核心的特性外,NATS Streaming提供了以下一些功能特性

## 加強消息協議
NATS Streaming使用Google緩衝區實現自己的增强型消息格式 這些消息通過二進位數據流在NATS核心平台進行傳遞, 因此不需要改變NATS的基本協議. NATS Streaming信息包含以下字段:
- 序列:一個全局順序序列號為主題的通道
- 主題:是NATS Streaming 交付對象
- 答覆内容:對應"reply-to"對應的對象内容
- 數據:真是數據内容
- 時間戳記:接收的時間戳記, 單位是奈秒(ns)
- 重複發送:標誌這調數據是否需要服務再次發送
- CRC32:一個循環冗餘數據校驗選項, 在數據存儲和數據通訊領域裡, 為了保證數據的正確性所採用的檢查手段, 這裡使用的是IEEE CRC32算法

## 消息/事件的持久性
　　NATS Streaming提供了可配置的消息持久化, 持久目的地可以為内存或者文件. 另外, 對應的存儲子系統使用了一個公共接口允許我們開發自己自定義實現来持久化對應的消息

## 至少一次的發送
　　NATS Streaming提供了發布者和server之間的消息確認(發布操作)和訂閱者和server之間的消息確認(確認消息發送). 其中消息被保存在server端内存或者輔助存儲(或其他外部存儲器)用来為需要重新接受消息的訂閱者並重發消息. 

## 發布者發送速率限定
　　NATS Streaming提供了一個連接選項叫MaxPubAcksInFlight, 它能有效的限制一個發布者可能随意的在任何时候發送的未被確認的消息. 當達到這個配置的最大數量時, 異步發送調用接口將會被阻塞, 直到未確認消息降到指定數量之下. 

## 每個訂閱者的速率匹配or限制
　　NATS Streaming運行指定的訂閱中設置一個參數為MaxInFlight, 它用來指定已確認但未消費的最大數據量, 當達到這個限制时, NATS Streaming將暫停發送消息給訂閱者, 直到未確認的數據量小於設定的量為止

## 以主題重發的歷史數據
　　新訂閱的可以在已經存儲起來的訂閱的主題頻道指定起始位置消息流. 通過使用這個選項,消息就可以開始發送傳遞了:
- 訂閱的主題存儲的最早的信息
- 與當前訂閱主題之前的最近存儲的數據, 這通常被認為是"最後的值"或"初值"對應的緩存
- 一個以奈秒為基准的日期or時間
- 一個歷史的起始位置相對當前服務的日期or時間, 例如：最後30秒
- 一個特定的消息序列號

## 持久訂閱
　　訂閱也可以指定一個“持久化的名稱”可以在客户端重新啟動時不受影響. 持久訂閱會使得對應服務跟踪客户端最後確認消息的序列號和持久名稱. 當這個客户端重啟或者重新訂閱的时候, 使用相同的客户端ID和持久化的名稱, 對應的服務將會從最早的未被確認的消息處恢複

## 安裝服務
指令:
```cmd
安裝
go get github.com/nats-io/nats-streaming-server
啟動服務
./nats-streaming-server
```

### 參數說明

```
--seq <seqno>           從指定的序列號開始
--all                   接受所有發送的數據
--last                  從上一次最後一次發的那比訊息開始
--since <duration>      從當前往前的時間段落開始接收(例如:1s, 1hr https://golang.org/pkg/time/#ParseDuration)
--durable <name>        永久訂閱者的名稱
--unsubscribe           退出出時解除永久訂閱
```


#### 參考文件
- https://docs.nats.io/nats-concepts/intro
- https://www.cnblogs.com/liang1101/p/6679180.html
- https://www.jishuwen.com/d/2EgL/zh-hk
- https://cloud.tencent.com/developer/article/1157816
- https://blog.csdn.net/chszs/article/details/50996484
- https://cloud.tencent.com/developer/article/1506392
- https://www.jianshu.com/p/5524dbde7204
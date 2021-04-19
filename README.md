# go-liziwei01-appui

在/go-liziwei01-appui/目录下使用
go run main.go
编译并运行APP

域名和接口设置在/go-liziwei01-appui/httpapi package下
目前ip:port设定为
localhost:8080

论文访问接口为/paperList
启动APP后使用以下url访问服务
http://localhost:8086/paperList?

get参数支持
pageIndex       默认1
pageLength      默认10
authors         默认空
startTime       默认当天0:00
endTime         默认当天24:00

搜索示例：
http://localhost:8086/paperList?authors=liziwei

library/env,library/conf内容参考了gdp
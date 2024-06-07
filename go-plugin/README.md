# golang plugin

## 注意

1.不支持 Windows 系统
2.不支持 调试模式
3.不支持 go-delve

## linux
```shell
mkdir plugins
go build -buildmode=plugin -o ./plugins ./extensions/chinese
go build -buildmode=plugin -o ./plugins ./extensions/english
go build -o main ./primary

./main

2024/06/07 17:48:40 INFO plugin lookup success pluginName=chinese.so
2024/06/07 17:48:40 INFO message from ChineseGreeter.SayHello
2024/06/07 17:48:40 INFO run SayHello by plugin result="Chinese : 你好 tom ... !"

2024/06/07 17:48:40 INFO plugin lookup success pluginName=english.so
2024/06/07 17:48:40 INFO message from EnglishGreeter.SayHello
2024/06/07 17:48:40 INFO run SayHello by plugin result="English : Hello tom ... !"

```

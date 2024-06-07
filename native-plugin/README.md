# dynamic library plugin

## 注意 

## linux
```shell
mkdir plugins

go build -buildmode=c-shared -o ./primary/libplugins.dll ./extensions/chinese

go build -buildmode=c-shared -o ./primary/libplugins.dll ./extensions/english

go build -o main ./primary

./main

2024/06/07 17:48:40 INFO plugin lookup success pluginName=chinese.so
2024/06/07 17:48:40 INFO message from ChineseGreeter.SayHello
2024/06/07 17:48:40 INFO run SayHello by plugin result="Chinese : 你好 tom ... !"

2024/06/07 17:48:40 INFO plugin lookup success pluginName=english.so
2024/06/07 17:48:40 INFO message from EnglishGreeter.SayHello
2024/06/07 17:48:40 INFO run SayHello by plugin result="English : Hello tom ... !"

```

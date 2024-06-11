# dynamic library plugin

## 注意 

## windows
```shell
go build -buildmode=c-shared -o ./libplugins.dll ./extensions/chinese
go build -o main.exe .
.\main.exe

2024/06/07 21:58:49 INFO message from ChineseGreeter.SayHello name=tom
2024/06/07 21:58:49 INFO SayHello result="Chinese : 你好 tom ... !"


go build -buildmode=c-shared -o ./libplugins.dll ./extensions/english
.\main.exe

2024/06/07 21:58:18 INFO message from EnglishGreeter.SayHello name=tom
2024/06/07 21:58:18 INFO SayHello result="English : Hello tom ... !"

```

## linux
```shell
# 将动态库环境变量设置为当前工作环境(不然找到不so文件)
export LD_LIBRARY_PATH=$GOPATH/src/go-plugin-examples/native-plugin

go build -buildmode=c-shared -o ./libplugins.so ./extensions/chinese
go build -o main .
./main

2024/06/07 21:58:49 INFO message from ChineseGreeter.SayHello name=tom
2024/06/07 21:58:49 INFO SayHello result="Chinese : 你好 tom ... !"


go build -buildmode=c-shared -o ./libplugins.so ./extensions/english
./main

2024/06/07 21:58:18 INFO message from EnglishGreeter.SayHello name=tom
2024/06/07 21:58:18 INFO SayHello result="English : Hello tom ... !"

```

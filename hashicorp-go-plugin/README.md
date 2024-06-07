# hashicorp-go-plugin

```shell
mkdir plugins
go build -o ./plugins ./extensions/chinese
go build -o ./plugins ./extensions/english
go build -o main ./primary
```

## windows
```shell
.\main.exe -plugin .\plugins\chinese.exe
2024-06-07T17:25:43.368+0800 [INFO]  chinese.exe: message from ChineseGreeter.SayHello: @module="Chinese Plugin" timestamp=2024-06-07T17:25:43.368+0800
2024/06/07 17:25:43 Chinese : 你好 tom ... !
2024-06-07T17:25:43.372+0800 [INFO]  plugin process exited: plugin=.\plugins\chinese.exe id=18908

.\main.exe -plugin .\plugins\english.exe
2024-06-07T17:25:51.866+0800 [INFO]  english.exe: message from EnglishGreeter.SayHello: @module="English Plugin" timestamp=2024-06-07T17:25:51.866+0800
2024/06/07 17:25:51 English : Hello tom ... !
2024-06-07T17:25:51.871+0800 [INFO]  plugin process exited: plugin=.\plugins\english.exe id=21712
```

## linux
```shell
./main -plugin ./plugins/chinese
2024-06-07T17:25:43.368+0800 [INFO]  chinese: message from ChineseGreeter.SayHello: @module="Chinese Plugin" timestamp=2024-06-07T17:25:43.368+0800
2024/06/07 17:25:43 Chinese : 你好 tom ... !
2024-06-07T17:25:43.372+0800 [INFO]  plugin process exited: plugin=.\plugins\chinese.exe id=18908

./main -plugin ./plugins/english
2024-06-07T17:25:51.866+0800 [INFO]  english: message from EnglishGreeter.SayHello: @module="English Plugin" timestamp=2024-06-07T17:25:51.866+0800
2024/06/07 17:25:51 English : Hello tom ... !
2024-06-07T17:25:51.871+0800 [INFO]  plugin process exited: plugin=.\plugins\english.exe id=21712
```

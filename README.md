# golang plugin example

| Name                | Linux              | Windows            | MacOS              | Other              | Details                                                      |
| ------------------- | ------------------ | ------------------ | ------------------ | ------------------ | ------------------------------------------------------------ |
| go-plugin           | :heavy_check_mark: | :x:                | :heavy_check_mark: | :heavy_check_mark: | 使用 go 1.8 原生支持的plugin的方式。<br />但是不支持windows系统 和 go-delve/delve 调试插件<br />必须安装C语言编译器。 |
| hashicorp-go-plugin | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | 使用 https://github.com/hashicorp/go-plugin <br />允许一个子进程，与子进程之间使用RPC通信。<br />无需 安装 C语言编译器 |
| native-plugin       | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | 将go库编译成动态库，如: dll,so。<br />然后以cgo的方式调用。需要会C语言且必须安装C语言编译器 |
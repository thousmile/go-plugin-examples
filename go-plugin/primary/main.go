package main

import (
	"flag"
	"fmt"
	"go-plugin/shared"
	"log/slog"
	"os"
	"plugin"
	"strings"
)

func main() {
	var pluginsDir string
	flag.StringVar(&pluginsDir, "plugins", "plugins", "path to plugins")
	flag.Parse()
	// 读取 plugins 目录下 所有的 .so 文件
	plugins, err := os.ReadDir(pluginsDir)
	if err != nil {
		slog.Error("error reading plugins directory: ", slog.Any("error", err))
		return
	}
	for _, entry := range plugins {
		// 判断文件 后缀名是否为 .so
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".so") {
			pluginPath := fmt.Sprintf("%s/%s", pluginsDir, entry.Name())
			// 打开 .so 文件
			p, err := plugin.Open(pluginPath)
			if err != nil {
				slog.Error("error opening plugin: ", slog.Any("error", err))
			} else {
				// 符号解析，获取插件中的 Impl 变量。必须是 插件中 定义 的变量名称
				sym, err := p.Lookup("Impl")
				if err != nil {
					slog.Error("error looking up plugin: ", slog.Any("error", err))
				} else {
					slog.Info("plugin lookup success", slog.String("pluginName", entry.Name()))
					// 使用反射 将插件 转成 接口
					impl := sym.(shared.Greeter)
					// 调用 接口 的方法
					resp := impl.SayHello("tom")
					slog.Info("run SayHello by plugin", slog.Any("result", resp))
				}
			}
			println()
		}
	}
}

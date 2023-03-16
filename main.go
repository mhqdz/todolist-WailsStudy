package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

// go注解 读取 frontend/dist 的所有内容
//
//go:embed all:frontend/dist
var assets embed.FS

// main 程序入口 调用了wails.Run() 定义了一些窗口选项 并指定了那些App可以暴露
func main() {
	// Create an instance of the app structure
	app := NewApp()

	// wails.Run()
	// Create application with options
	err := wails.Run(&options.App{
		Title:         "todoList",
		Width:         400,
		Height:        600,
		MaxHeight:     3000,
		DisableResize: true,
		// 无窗口 关闭
		Frameless: true,
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			// BackdropType: windows.Acrylic,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		OnStartup:        app.startup,

		// Bind 方法绑定 它会检查 Bind 字段中列出的结构体实例
		// 确定哪些函数是公开的(公开以大写字母开头的函数)
		// 并生成前端可以调用的这些方法的 JavaScript 版本
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

# README

## 关于
项目使用了官方的Wails Vue模板。您可以通过编辑wills .json来配置项目

项目配置信息[wails官方文档](https://wails.io/zh-Hans/docs/reference/project-config/).

自己写的稍微更详细一点的[项目代码解析](mhqdz.cn/posts/wails/todolist(wailsStudy))

## 构建
项目构建依赖于[wails](https://wails.io).

请在项目目录中使用`wails build`来编译它 编译结果在 `build/bin`目录下
```powershell
wails build
```
或使用`wails dev`以开发模式运行它 此时它会同时在 http://localhost:34115 开启一个服务
```powershell
wails dew
```
## Todo
- 整体样式总是不能如意 如果您擅长ui设计或者前端 请务必给它一个更合适的样式
- 窗口位置 希望它能出现在一个更合适的位置 而不是屏幕中心 但我不知道怎么做
- 拖拽延迟 我知道并不舒服 但不知如何解决
- 它目前对触摸屏并不友好 支持一些莫名其妙的功能 比如双指放大和横向滚动
- 数据写在 `~/mhTodolist`有些电脑有磁盘保护 哪怕设置666的文件权限依旧无法写入 不知道该怎么解决 所以写了一个去除只读权限的脚本 不过非常丑陋 如果有更好的解决办法就好了
- 开机启动 也是用了一个脚本实现 不希望它出现在面板上 又不希望它在使用后就会开机启动 所以折中的用了一个脚本 如果ui被重新设计 更希望它被放在可交互的某个角落
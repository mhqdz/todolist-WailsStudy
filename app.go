package main

import (
	"bufio"
	"context"
	"io"
	"math/rand"
	"os"
	"os/user"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// ctx 上下文
// Msgs 存储todos的数组
// say 存储
// App struct
type App struct {
	ctx  context.Context
	Msgs []string
	say  string
}

// NewApp App struct的构造函数
func NewApp() *App {
	app := &App{}
	app.Read()
	return app
}

// KillSelf 关闭程序
func (a *App) KillSelf() {
	os.Exit(0)
}

// MinWindow 窗口最小化
func (a *App) MinWindow() {
	runtime.WindowMinimise(a.ctx)
}

// 下边的那一句话
// RedomSay 随机获取一个合适的语句(say) 写入到app.say是为了防止随机到同一句话
func (a *App) RedomSay() string {
	says := []string{"如我所愿,一切皆好", "不可预见的未来是诗 写在纸上的未来是理想"}

	now := time.Now()
	mounth := now.Month()
	hour := now.Hour()

	// 春
	if mounth == 3 || mounth == 4 || mounth == 5 {
		says = append(says, "正人间四月 何不绽放如花")
		says = append(says, "山中何事?松花酿酒,春水煎茶")
	}

	// 夏
	if mounth == 6 || mounth == 7 || mounth == 8 {
		says = append(says, "生如夏花之绚烂")
		says = append(says, "今天会下雨吗")
	}

	// 秋
	if mounth == 9 || mounth == 10 || mounth == 11 {
		says = append(says, "空山新雨后,天气晚来秋")
	}

	// 冬
	if mounth == 12 || mounth == 1 || mounth == 2 {
		if (hour > 20 && hour <= 23) || (hour > 0 && hour < 6) {
			says = append(says, "秋月扬明晖，冬岭秀寒松")
		}
	}

	// 凌晨
	if hour >= 0 && hour < 6 {
		says = append(says, "永恒的月光和永动的齿轮 那是永远的良药")
		says = append(says, "在白天的残骸上 人们徒然睁着双眼")
		says = append(says, "白昼之光,岂知夜之深")
	}

	// 早晨
	if hour > 5 && hour < 8 {
		says = append(says, "洗漱之后 开始今天的安排吧")
		says = append(says, "日出雾露馀，青松如膏沐")
	}

	// 上午
	if hour > 7 && hour < 11 {
		says = append(says, "梦想和苦恼都已过去 这是新的一天")
	}

	// 中午
	if hour > 11 && hour < 14 {
		says = append(says, "午睡或者看会书 无论如何 这暧昧的时间属于你")
	}

	// 下午
	if hour > 13 && hour < 17 {
		says = append(says, "久坐是反对神圣精神的罪恶 感到烦闷不妨出去走走")
		says = append(says, "惟其不可能,才值得相信")
		if hour == 14 {
			says = append(says, "来杯下午茶吧")
		}
	}

	// 晚饭
	if hour > 16 && hour < 21 {
		says = append(says, "不要忘了晚饭哦")
		says = append(says, "看到密涅瓦的猫头鹰了吗")
	}

	// 晚上
	if hour > 20 && hour <= 23 {
		says = append(says, "保持清洁 洗漱是必要的")
		says = append(says, "给今天的工作收个尾吧")
	}

	sayLen := len(says)

	if sayLen == 0 {
		a.say = "如我所愿,一切皆好"
		return a.say
	}

	if sayLen == 1 {
		a.say = says[0]
		return a.say
	}

	var i int32
	for {
		i2 := rand.Int31n(int32(sayLen))
		if says[i2] != a.say {
			i = i2
			break
		}
	}

	a.say = says[i]
	return a.say
}

// startup 相当于init函数 在启动时执行
// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// WindowSetPosition 设置启动位置 而不是总在屏幕中间
	// 可以在上面任意会被执行的函数里插入一个 fmt.Println(runtime.WindowGetPosition(a.ctx)) 用wails dev运行 来打印当前坐标 选择合适的坐标 然后在这里设置启动时就到那个位置
	// 当然如果您愿意的话 可以用某个文件 记录下关闭时的位置 然后读取 改用那个位置
	// runtime.WindowSetPosition(ctx, -1051, 456)
}

// SetH 设置窗口高度为h h具体的值在前端计算获得
func (a *App) SetH(h int) {
	// fmt.Println(runtime.WindowGetSize(a.ctx))
	// fmt.Println(h)
	runtime.WindowSetSize(a.ctx, 400, h)
}

// Add 添加一个安排
func (a *App) Add(todo string) {
	if todo == "" {
		return
	}
	a.Read()
	a.Msgs = append(a.Msgs, todo)
	a.Write()
	// fmt.Println(runtime.WindowGetPosition(a.ctx))
	// fmt.Println(a.Msgs)
}

// GetLen 获取当前
func (a *App) GetLen() int {
	return len(a.Msgs)
}

// GetMsgs 返回存有 todos的数组
func (a *App) GetMsgs() []string {
	a.Read()
	return a.Msgs
}

// Write 把a.msgs写入到用户目录下的mhtodolist
func (a *App) Write() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(user.HomeDir+"\\mhtodolist", os.O_WRONLY|os.O_CREATE, 0666)

	// 在写入前把文件清空 避免源文件比写入得到文件大的情况
	file.Truncate(0)

	// file, err := os.OpenFile("/static/mhtodolist",os.O_WRONLY|os.O_CREATE,0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, j := range a.Msgs {
		if j[len(j)-1] == '\n' {
			writer.WriteString(j)
			continue
		}
		writer.WriteString(j + "\n")
	}
	writer.Flush()
}

// 没有用上 毕竟上面的GetMsgs已经把整个数组都拿去了 没必要来后端拿
// GetMsg 根据index获取一条 todo
func (a *App) GetMsg(index int) string {
	return a.Msgs[index]
}

// Del 根据index删除一个todo
func (a *App) Del(index int) {
	a.Msgs = append(a.Msgs[:index], a.Msgs[index+1:]...)
	a.Write()
}

// Set 编辑第index条信息
func (a *App) Set(index int, todo string) {
	if todo == "" {
		return
	}
	a.Read()
	a.Msgs[index] = todo
	a.Write()
	// fmt.Println(a.Msgs)
}

// Read 把mhtodolist内容读取到a.Msgs里
func (a *App) Read() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(user.HomeDir+"\\mhtodolist", os.O_RDONLY|os.O_CREATE, 111)
	// err = nil
	defer file.Close()

	reder := bufio.NewReader(file)
	arr := []string{}
	for {
		line, err := reder.ReadString('\n')
		// err = nil
		if err == io.EOF {
			if len(line) != 0 {
				arr = append(arr, line)
			}
			break
		}
		if err != nil {
			panic("读取失败:" + err.Error())
		}
		arr = append(arr, line)
	}
	a.Msgs = arr
}

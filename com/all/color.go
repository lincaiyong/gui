package all

import "fmt"

func color(r, g, b int) string {
	return fmt.Sprintf("'#%02x%02x%02x'", r, g, b)
}

const (
	ColorBlack   = "'black'"   // 黑色
	ColorSilver  = "'silver'"  // 银色
	ColorGray    = "'gray'"    // 灰色
	ColorWhite   = "'white'"   // 白色
	ColorMaroon  = "'maroon'"  // 栗色
	ColorRed     = "'red'"     // 红色
	ColorPurple  = "'purple'"  // 紫色
	ColorMagenta = "'magenta'" // 洋红色
	ColorGreen   = "'green'"   // 绿色
	ColorLime    = "'lime'"    // 酸橙色
	ColorOlive   = "'olive'"   // 橄榄色
	ColorYellow  = "'yellow'"  // 黄色
	ColorNavy    = "'navy'"    // 海军蓝
	ColorBlue    = "'blue'"    // 蓝色
	ColorTeal    = "'teal'"    // 蓝绿色
	ColorCyan    = "'cyan'"    // 水绿色/青色
)

var (
	ColorGray247 = color(247, 248, 250)
	ColorGray235 = color(235, 236, 240)
	ColorGray201 = color(201, 204, 213)
)

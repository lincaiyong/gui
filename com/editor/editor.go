package editor

import (
	"fmt"
	"github.com/lincaiyong/gui/com"
)

const (
	LangGo         = "'go'"
	LangJava       = "'java'"
	LangJavascript = "'javascript'"
	LangPython     = "'python'"
	LangPhp        = "'php'"
)

func Editor() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret)
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	showLineNo             com.Property `default:"true"`
	onCursorPositionChange com.Property `default:"undefined"`

	onCreated   com.Method
	onUpdated   com.Method
	_destroy    com.Method
	setValue    com.Method
	getValue    com.Method
	setLanguage com.Method
}

func (c *Component) ShowLineNo(b bool) *Component {
	c.SetProp("showLineNo", fmt.Sprintf("%v", b))
	return c
}

func (c *Component) OnCursorPositionChange(s string) *Component {
	c.SetProp("onCursorPositionChange", s)
	return c
}

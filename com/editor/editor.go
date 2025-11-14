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
	value                  com.Property `default:"''"`
	language               com.Property `default:"'go'"`
	showLineNo             com.Property `default:"true"`
	onCursorPositionChange com.Property `default:"undefined"`

	onCreated com.Method
	onUpdated com.Method
	_destroy  com.Method
}

func (c *Component) Value(s string) *Component {
	c.SetProp("value", s)
	return c
}

func (c *Component) Language(s string) *Component {
	c.SetProp("language", s)
	return c
}

func (c *Component) ShowLineNo(b bool) *Component {
	c.SetProp("showLineNo", fmt.Sprintf("%v", b))
	return c
}

func (c *Component) OnCursorPositionChange(s string) *Component {
	c.SetProp("onCursorPositionChange", s)
	return c
}

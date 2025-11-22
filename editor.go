package gui

import (
	"fmt"
)

func NewEditorOpt() *EditorOpt {
	ret := &EditorOpt{}
	ret.BaseOpt = NewBaseOpt[EditorOpt](ret)
	ret.Value("''").Language("'go'").ShowLineNo(true).OnCursorPositionChange("null")
	return ret
}

type EditorOpt struct {
	*BaseOpt[EditorOpt]
}

func (o *EditorOpt) Value(s string) *EditorOpt {
	o.SetProperty("value", s)
	return o
}

func (o *EditorOpt) Language(s string) *EditorOpt {
	o.SetProperty("language", s)
	return o
}

func (o *EditorOpt) ShowLineNo(b bool) *EditorOpt {
	o.SetProperty("showLineNo", fmt.Sprintf("%v", b))
	return o
}

func (o *EditorOpt) OnCursorPositionChange(s string) *EditorOpt {
	o.SetStaticProperty("onCursorPositionChange", s)
	return o
}

func Editor(opt *EditorOpt) *Element {
	ret := NewElement(ElementTypeEditor, ElementTagDiv)
	opt.OnCreated("editor_handleCreated").OnUpdated("editor_handleUpdated").OnDestroy("editor_handleDestroy").Init(ret)
	return ret
}

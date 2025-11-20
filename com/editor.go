package com

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
	o.SetProperty("onCursorPositionChange", s)
	return o
}

func Editor(opt EditorOpt) *Element {
	ret := NewElement(ElementTypeEditor, ElementTagDiv)
	ret.SetMethod("onCreated", `function() {
    let lineNumbers = 'on';
    if (!this.showLineNo) {
        lineNumbers = 'off';
    }
    const value = this.value;
    const language = this.language;
    const options = {
        value,
        language,
        theme: 'vs',
        automaticLayout: true,
        lineNumbers,
        minimap: {
            enabled: false,
        },
        readOnly: true,
        // fontFamily: '',
        // glyphMargin: false,
        // suggestOnTriggerCharacters: false,
    };
    this._editor = monaco.editor.create(this.ref, options);
    this._editor.onDidChangeCursorPosition((e) => {
        this.onCursorPositionChange?.(e.position.lineNumber, e.position.column);
    });
    this._editor.onDidChangeModelContent(() => {
        this.value = this._editor.getValue();
    });
}`).
		SetMethod("onUpdated", `function(k, v) {
    if (!this._editor) {
        return;
    }
    switch (k) {
        case 'value':
            this._editor.setValue(v);
            break;
        case 'language':
            monaco.editor.setModelLanguage(this._editor.getModel(), v);
            break;
    }
}`).
		SetMethod("onDestroy", `function() {
    this._editor.dispose();
}`)
	opt.Init(ret)
	return ret
}

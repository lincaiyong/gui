package com

func Compare() *Element {
	ret := NewElement("compare", "div")
	ret.SetMethod("onCreated", `function() {
    const leftModel = monaco.editor.createModel('原始文本', 'text/plain');
    const rightModel = monaco.editor.createModel('修改后的文本', 'text/plain');

    this._editor = monaco.editor.createDiffEditor(this.ref, {
        automaticLayout: true,
    });
    this._editor.setModel({
        original: leftModel,
        modified: rightModel
    });
}`).SetMethod("onDestroy", `function() {
    this._editor.dispose();
}`)
	return ret
}

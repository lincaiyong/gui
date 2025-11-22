package gui

func Compare(opt *Opt) *Element {
	ret := NewElement(ElementTypeCompare, ElementTagDiv)
	ret.SetMethod("handleCreated", `function() {
    const leftModel = monaco.editor.createModel('原始文本', 'text/plain');
    const rightModel = monaco.editor.createModel('修改后的文本', 'text/plain');

    this._editor = monaco.editor.createDiffEditor(this.ref, {
        automaticLayout: true,
    });
    this._editor.setModel({
        original: leftModel,
        modified: rightModel
    });
}`).SetMethod("handleDestroy", `function() {
    this._editor.dispose();
}`)
	opt.OnCreated(".handleCreated").OnDestroy(".handleDestroy").Init(ret)
	return ret
}

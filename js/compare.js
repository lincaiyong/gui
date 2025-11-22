function compare_handleCreated() {
    const leftModel = monaco.editor.createModel('原始文本', 'text/plain');
    const rightModel = monaco.editor.createModel('修改后的文本', 'text/plain');

    this._editor = monaco.editor.createDiffEditor(this.ref, {
        automaticLayout: true,
    });
    this._editor.setModel({
        original: leftModel,
        modified: rightModel
    });
}

function compare_handleDestroy() {
    this._editor.dispose();
}
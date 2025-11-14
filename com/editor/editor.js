function _destroy() {
    this._editor.dispose();
    super._destroy();
}

function onCreated() {
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
}

function onUpdated(k, v) {
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
}
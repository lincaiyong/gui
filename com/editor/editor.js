function _destroy() {
    this._editor.dispose();
    super._destroy();
}

function setValue(v) {
    this._editor.setValue(v);
}

function getValue() {
    return this._editor.getModel().getValue();
}

function setLanguage(v) {
    monaco.editor.setModelLanguage(this._editor.getModel(), v);
}

function onCreated() {
    let lineNumbers = 'on';
    if (!this.showLineNo) {
        lineNumbers = 'off';
    }
    const options = {
        value: '',
        language: '',
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
}

function onUpdated(k, v) {
    if (!this._editor) {
        return;
    }
    switch (k) {
    case 'showLineNo':
        if (v) {
            this._editor.updateOptions({ lineNumbers: 'on' });
        } else {
            this._editor.updateOptions({ lineNumbers: 'off' });
        }
    }
}
function _destroy() {
    this._editor.dispose();
    super._destroy();
}

function setValue(v) {
    this._editor.setValue(v);
}

function setLanguage(v) {
    monaco.editor.setModelLanguage(this._editor.getModel(), v);
}

function onCreated() {
    const options = {
        value: '',
        language: '',
        theme: 'vs',
        automaticLayout: true,
        lineNumbers: 'on', // 'off'
        minimap: {
            enabled: false,
        },
        readOnly: true,
        // fontFamily: '',
        // glyphMargin: false,
        // suggestOnTriggerCharacters: false,
    };
    this._editor = monaco.editor.create(this.ref, options);
}
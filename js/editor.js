function editor_handleCreated(ele) {
    let lineNumbers = 'on';
    if (!ele.showLineNo) {
        lineNumbers = 'off';
    }
    const value = ele.value;
    const language = ele.language;
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
    ele._editor = monaco.editor.create(ele.ref, options);
    ele._editor.onDidChangeCursorPosition((e) => {
        ele.onCursorPositionChange?.(e.position.lineNumber, e.position.column);
    });
    ele._editor.onDidChangeModelContent(() => {
        ele.value = ele._editor.getValue();
    });
}

function editor_handleUpdated(ele, k, v) {
    if (!ele._editor) {
        return;
    }
    switch (k) {
        case 'value':
            ele._editor.setValue(v);
            break;
        case 'language':
            monaco.editor.setModelLanguage(ele._editor.getModel(), v);
            break;
    }
}

function editor_handleDestroy(ele) {
    ele._editor.dispose();
}
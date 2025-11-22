function compare_handleCreated(ele) {
    const leftModel = monaco.editor.createModel('原始文本', 'text/plain');
    const rightModel = monaco.editor.createModel('修改后的文本', 'text/plain');

    ele._editor = monaco.editor.createDiffEditor(ele.ref, {
        automaticLayout: true,
    });
    ele._editor.setModel({
        original: leftModel,
        modified: rightModel
    });
}

function compare_handleDestroy(ele) {
    ele._editor.dispose();
}
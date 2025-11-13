function log(v) {
    go.main.App.Log(v);
}
function queryDefinition(file, lineIdx, charIdx) {
    return go.main.App.QueryDefinition(file, lineIdx, charIdx);
}
function openProject() {
    return go.main.App.OpenProject();
}
function openFile(path) {
    return go.main.App.OpenFile(path);
}

function onClickTreeItem(itemEle) {
    Root.log('click: ' + JSON.stringify(itemEle.data));
    if (itemEle.data.leaf) {
        g.root.currentFile = itemEle.data.key;
        const relPath = itemEle.data.key;
        Root.openFile(g.root.sourceRoot + '/' + relPath).then(v => {
            g.root.editorEle.setValue(v);
        });
    }
}

function onOpenProject() {
    Root.openProject().then(s => {
        const obj = JSON.parse(s)
        Root.log(obj);
        g.root.sourceRoot = obj.folder;
        g.root.treeEle.items = obj.files;
    });
}

function onStartup() {
    setTimeout(function() {
        g.root.treeItems = ['com/test.go', 'res.go'];
    });
}

function getColIdx(code, lineNo, charNo) {
    const lines = code.split('\n');
    if (lineNo < 1 || lineNo > lines.length) {
        return 0;
    }
    const lineContent = lines[lineNo - 1];
    if (charNo < 1 || charNo > lineContent.length) {
        return 0;
    }
    let colIdx = 0;
    for (let i = 0; i < charNo; i++) {
        if (lineContent[i] === '\t') {
            colIdx += 4 - ((colIdx - 1) % 4);
        } else {
            colIdx += 1;
        }
    }
    return colIdx;
}

function onCursorPositionChange(lineNo, charNo) {
    Root.log(`cursor: ${lineNo} ${charNo}`);
    const content = g.root.editorEle.getValue();
    const colIdx = Root.getColIdx(content, lineNo, charNo);
    Root.queryDefinition(g.root.currentFile, lineNo-1, colIdx).then(v => {
        Root.log(v);
        g.root.outputEle.setValue(v);
    })
}
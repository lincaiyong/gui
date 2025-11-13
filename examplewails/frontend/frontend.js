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
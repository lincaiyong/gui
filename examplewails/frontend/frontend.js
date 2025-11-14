function log(v) {
    go.main.App.Log(v);
}

function onClickTreeItem(itemEle) {
    Root.log('click: ' + JSON.stringify(itemEle.data));
    if (itemEle.data.leaf) {
        g.root.currentFile = itemEle.data.key;
        const relPath = itemEle.data.key;
        go.main.App.OpenFile(g.root.sourceRoot + '/' + relPath).then(v => {
            g.root.currentFileContent = v;
            if (relPath.endsWith('.go')) {
                g.root.currentFileLanguage = 'go';
            }
        });
    }
}

function openProject(folder) {
    go.main.App.OpenProject(folder).then(s => {
        const obj = JSON.parse(s)
        Root.log(obj);
        g.root.files = obj.files;
        g.root.sourceRoot = obj.folder;
        g.root.projectName = obj.folder.substring(obj.folder.lastIndexOf('/')+1);
    });
}

function onOpenDirectory() {
    go.main.App.OpenDirectory().then(folder => {
        if (folder) {
            Root.openProject(folder);
        }
    });
}

function OnCreated() {
    setTimeout(function() {
        Root.openProject('/Users/andy/Code/lincaiyong/gui');
    });
}

function onCursorPositionChange(lineNo, charNo) {
    Root.log(`cursor: ${lineNo} ${charNo}`);
    go.main.App.QueryDefinition(`${g.root.sourceRoot}/${g.root.currentFile}`, lineNo-1, charNo-1).then(s => {
        Root.log(`definition: ${s}`);
        const targets = JSON.parse(s);
        if (targets && targets.length === 1) {
            const target = targets[0];
            let file = target.file;
            if (target.file.includes(g.root.sourceRoot)) {
               file = file.substring(g.root.sourceRoot.length+1);
            }
            g.root.outputText = `${file}#L${target.lineIdx +1}`;
        }
    })
}
function _clearUndo() {
    const {page} = this.data;
    page._undoStack = [];
    page._redoIndex = 0;
}

function _pushUndo(command, text, range) {
    const {page} = this.data;
    const trans = {
        command: command,
        text: text,
        range: range,
    };
    if (!page._freezeUndo) {
        if (page._redoIndex < page._undoStack.length) {
            page._undoStack.splice(page._redoIndex);
        }
        page._undoStack.push(trans);
        page._redoIndex = page._undoStack.length;
    }
    return trans;
}

function freezeUndo(fun) {
    const {page} = this.data;
    page._freezeUndo = true;
    fun();
    page._freezeUndo = false;
}

function executeUndo() {
    const {page, cursor} = this.data;
    if (page._redoIndex === 0) {
        return;
    }
    page._redoIndex = page._redoIndex - 1;
    const trans = page._undoStack[page._redoIndex];
    // console.log('undo', trans);
    if (trans.command === webapp.constant.deleteTextCommand) {
        this.clearAnchor();
        this.moveCursorToPosition(trans.range.start);
        this.freezeUndo(() => this.doInsertText(trans.text));
        const cursorPos = this.copyPosition(cursor.position);
        this.moveCursorToPosition(trans.range.start);
        this.createAnchor();
        this.moveCursorToPosition(cursorPos);
        this.updateSelection();
        this.refreshSelection();
    } else if (trans.command === webapp.constant.insertTextCommand) {
        this.moveCursorToPosition(trans.range.start);
        this.createAnchor();
        this.moveCursorToPosition(trans.range.end);
        this.freezeUndo(() => this.doDeleteText());
    }
}

function executeRedo() {
    const {page, cursor} = this.data;
    if (page._redoIndex === page._undoStack.length) {
        return;
    }
    const trans = page._undoStack[page._redoIndex];
    page._redoIndex = page._redoIndex + 1;
    // console.log('redo', trans);
    if (trans.command === webapp.constant.deleteTextCommand) {
        this.moveCursorToPosition(trans.range.start);
        this.createAnchor();
        this.moveCursorToPosition(trans.range.end);
        this.freezeUndo(() => this.doDeleteText());
    } else if (trans.command === webapp.constant.insertTextCommand) {
        this.clearAnchor();
        this.moveCursorToPosition(trans.range.start);
        this.freezeUndo(() => this.doInsertText(trans.text));
        const cursorPos = this.copyPosition(cursor.position);
        this.moveCursorToPosition(trans.range.start);
        this.createAnchor();
        this.moveCursorToPosition(cursorPos);
        this.updateSelection();
        this.refreshSelection();
    }
}
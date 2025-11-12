function doCreatePage(text) {
    this.createPageContent(text);
    this.refreshPage();
    this.clearAnchor();
    this.hideCursor();
    this.refreshCursor(false);
    this.updateSelection();
    this.refreshSelection();
    this.activeLineEle.v = 0;
    this.refreshHighlight([]);
}

function doInsertText(text) {
    if (this.hasVisibleRange()) {
        this.doDeleteText();
    }
    const newCursorPos = this.pageInsertText(text);
    this.refreshPage();
    this.moveCursorToPosition(newCursorPos);
    this.refreshCursor(true);
    this.clearAnchor();
    this.updateSelection();
    this.refreshSelection();
}

function doDeleteText() {
    // page
    const {newCursorPos, deletedText} = this.pageDeleteText();
    this.refreshPage();
    // cursor
    this.moveCursorToPosition(newCursorPos);
    this.refreshCursor(true);
    // selection
    this.clearAnchor();
    this.updateSelection();
    this.refreshSelection();
    return deletedText;
}

function doCopyText() {
    return this.pageGetText();
}

function doCutText() {
    return this.doDeleteText();
}

function doPasteText(text) {
    this.doInsertText(text);
}
function updateSelection() {
    const {anchor, page} = this.data;
    if (!anchor.valid) {
        return;
    }
    const items = [];
    const {start, end} = this.getRange();
    if (start.lineIndex === end.lineIndex) {
        if (end.colIndex > start.colIndex) {
            items.push({lineIndex: start.lineIndex, start: start.colIndex, end: end.colIndex});
        }
    } else {
        items.push({lineIndex: start.lineIndex, start: start.colIndex, end: -1});
        for (let i = start.lineIndex + 1; i < end.lineIndex; i++) {
            items.push({lineIndex: i, start: 0, end: -1});
        }
        if (end.colIndex > 0) {
            items.push({lineIndex: end.lineIndex, start: 0, end: end.colIndex});
        }
    }
    page.selectionItems = items;
}

function getRange() {
    const {anchor, cursor} = this.data;
    const anchorPos = this.copyPosition(anchor.position);
    const cursorPos = this.copyPosition(cursor.position);
    const ret = this.newRange(cursorPos, cursorPos);
    if (anchor.valid) {
        ret.end = anchorPos;
        if (anchorPos.lineIndex < cursorPos.lineIndex || (anchorPos.lineIndex === cursorPos.lineIndex && anchorPos.colIndex < cursorPos.colIndex)) {
            ret.start = anchorPos;
            ret.end = cursorPos;
        }
    }
    return ret;
}

function hasVisibleRange() {
    const {anchor, cursor} = this.data;
    return anchor.valid && !this.samePosition(anchor.position, cursor.position);
}

function refreshSelection() {
    const {anchor, page} = this.data;
    if (anchor.valid) {
        this.selectionEle.v = 1;
        this.selectionEle.items = page.selectionItems;
    } else {
        this.selectionEle.v = 0;
    }
}
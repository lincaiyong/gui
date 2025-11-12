function handleScrollTop(ele, val) {
    this.refreshCursor(false);
}

function handleScrollLeft(ele, val) {
    this.refreshCursor(false);
}

function handleMouseDown(ele, mouseDownEvent) {
    const {keyPressed, anchor, cursor} = this.data;
    const startCursorPos = this.copyPosition(cursor.position);
    if (!keyPressed.shift && anchor.valid) {
        this.clearAnchor();
    } else if (keyPressed.shift && !anchor.valid) {
        this.createAnchor();
    }
    this._showCursorOnMouseEvent(mouseDownEvent);
    const cancelMouseMoveListener = webapp.event.addListener(
        window,
        'mousemove',
        (ev) => {
            if (!anchor.valid && !this.samePosition(startCursorPos, cursor.position)) {
                this.createAnchor();
            }
            this._showCursorOnMouseEvent(ev);
        },
    );
    webapp.event.onceListener(
        window,
        'mouseup',
        (ev) => {
            this._showCursorOnMouseEvent(ev);
            cancelMouseMoveListener();
        },
    );
}

function getWordByCursor() {
    const {page, cursor} = this.data;
    const {lineIndex, colIndex} = cursor.position;
    const {items} = page;
    const ret = {left: undefined,  right: undefined};
    if (lineIndex < items.length) {
        this.checkWordItems(lineIndex);
        const words = page.items[lineIndex].wordItems;
        // TODO: 二分法
        for (let i = 0; i < words.length; i++) {
            const word = words[i];
            if (word.start <= colIndex && word.end >= colIndex) {
                ret.left = word;
                if (word.end === colIndex && i < words.length - 1) {
                    ret.right = words[i+1];
                }
                break;
            }
        }
    }
    return ret;
}

function handleDoubleClick(ele, ev) {
    const {left, right} = this.getWordByCursor(false);
    let word = left.isWordLike ? left : right;
    this.moveCursorToCol(word.start);
    this.createAnchor();
    this.moveCursorToCol(word.end);
    this.refreshCursor(true);
    this.updateSelection();
    this.refreshSelection();
}

function _showCursorOnMouseEvent(ev) {
    const rect = this.pageEle.ref.getBoundingClientRect();
    this.moveCursorToPixel(
        ev.clientX - rect.x,
        ev.clientY - rect.y - 9,
    );
    this.refreshCursor(true);
    this.updateSelection();
    this.refreshSelection();
}
function checkWordItems(lineIndex) {
    const {page} = this.data;
    const {lines, items} = page;
    if (lineIndex >= items.length || items[lineIndex].wordItems.length > 0) {
        return;
    }
    const text = lines[lineIndex].join('');
    items[lineIndex].wordItems = webapp.util.text2words(text);
}

function moveCursorToPixel(x, y) {
    const {page} = this.data;
    const lineItems = page.items;
    // real x y
    x = x + this.contentEle.scrollLeft;
    y = y + this.contentEle.scrollTop;
    let lineIndex = Math.round(y / 20);
    if (lineIndex < 0) {
        lineIndex = 0;
    }
    if (lineIndex >= page.lines.length) {
        lineIndex = page.lines.length - 1;
    }
    const {colItems} = lineItems[lineIndex];
    // TODO: 二分法
    let colIndex = colItems.length - 1;
    for (let i = 0; i < colItems.length; i++) {
        if (colItems[i].left + colItems[i].width / 2 > x) {
            colIndex = i;
            break;
        }
    }
    this.moveCursorToPosition(this.newPosition(lineIndex, colIndex));
}

function moveCursorLeftIgnoreRange() {
    const {page, cursor} = this.data;
    const {lineIndex, colIndex} = cursor.position;
    if (colIndex > 0) {
        this.moveCursorToCol(colIndex - 1);
    } else if (lineIndex > 0) {
        const col = page.lines[lineIndex - 1].length;
        this.moveCursorToPosition(this.newPosition(lineIndex - 1, col));
    }
}

function moveCursorLeft() {
    if (this.hasVisibleRange()) {
        const {start} = this.getRange();
        this.moveCursorToPosition(start);
    } else {
        this.moveCursorLeftIgnoreRange();
    }
}

function moveCursorLeftByLine() {
    this.moveCursorToCol(0);
}

function moveCursorLeftByWord() {
    const {page, cursor} = this.data;
    const {lines, items} = page;
    const {lineIndex, colIndex} = cursor.position;
    if (lines[lineIndex].length === 0 || colIndex === 0) {
        return;
    }
    this.checkWordItems(lineIndex);
    const words = items[lineIndex].wordItems;
    // TODO: 二分法
    for (let i = 0; i < words.length; i++) {
        if (words[i].start < colIndex && words[i].end >= colIndex) {
            this.moveCursorToCol(words[i].start);
            return;
        } else if (words[i].start === colIndex && i > 0) {
            this.moveCursorToCol(words[i-1].start);
            return;
        }
    }
}

function moveCursorRightIgnoreRange() {
    const {page, cursor} = this.data;
    const {lineIndex, colIndex} = cursor.position;
    const maxCol = page.lines[lineIndex].length;
    if (colIndex < maxCol) {
        this.moveCursorToCol(colIndex + 1);
    } else if (lineIndex < page.lines.length - 1) {
        this.moveCursorToPosition(this.newPosition(lineIndex + 1, 0));
    }
}

function moveCursorRight() {
    if (this.hasVisibleRange()) {
        const {end} = this.getRange();
        this.moveCursorToPosition(end);
    } else {
        this.moveCursorRightIgnoreRange();
    }
}

function moveCursorRightByLine() {
    const {page, cursor} = this.data;
    const {lineIndex} = cursor.position;
    this.moveCursorToCol(page.lines[lineIndex].length);
}

function moveCursorRightByWord() {
    const {page, cursor} = this.data;
    const {lineIndex, colIndex} = cursor.position;
    const {lines, items} = page;
    if (lines[lineIndex].length === 0) {
        return;
    }
    this.checkWordItems(lineIndex);
    const words = items[lineIndex].wordItems;
    // TODO: 二分法
    for (let i = 0; i < words.length; i++) {
        if (words[i].start <= colIndex && words[i].end > colIndex) {
            this.moveCursorToCol(words[i].end);
            return;
        } else if (words[i].end === colIndex && i < words.length - 1) {
            this.moveCursorToCol(words[i+1].end);
            return;
        }
    }
}

function moveCursorUp() {
    const {position} = this.data.cursor;
    this.moveCursorToLine(position.lineIndex - 1);
}

function moveCursorDown() {
    const {position} = this.data.cursor;
    this.moveCursorToLine(position.lineIndex + 1);
}

function moveCursorToLine(line) {
    const {page, cursor} = this.data;
    const lineItems = page.items;
    if (line >= lineItems.length) {
        line = lineItems.length - 1;
    }
    if (line < 0) {
        line = 0;
    }
    cursor.position.lineIndex = line;
    this.moveCursorToCol(cursor.position.colIndex);
}

function moveCursorToCol(col) {
    const {page, cursor} = this.data;
    const lineItems = page.items;
    const {colItems} = lineItems[cursor.position.lineIndex];
    if (col >= colItems.length) {
        col = colItems.length - 1;
    }
    if (col < 0) {
        col = 0;
    }
    cursor.position.colIndex = col;
    cursor.visible = true;
}

function moveCursorToPosition(pos) {
    this.moveCursorToLine(pos.lineIndex);
    this.moveCursorToCol(pos.colIndex);
}

function moveCursorToPageStart() {
    this.moveCursorToPosition(this.newPosition(0, 0));
}

function moveCursorToPageEnd() {
    const {page} = this.data;
    this.moveCursorToPosition(this.newPosition(
        page.lines.length - 1,
        page.lines[page.lines.length - 1].length,
    ));
}

function hideCursor() {
    this.data.cursor.visible = false;
    this.refreshCursor(false);
}

function refreshCursor(scrollPage) {
    const {page, cursor} = this.data;
    if (!cursor.visible) {
        const {lineIndex} = cursor.position;
        this.cursorEle.x = -2;
        this.activeLineEle.y = lineIndex * 20 - this.contentEle.scrollTop;
        return;
    }
    // 确保位置是合法的
    this.moveCursorToPosition(cursor.position);
    const {lineIndex, colIndex} = cursor.position;
    const lineItems = page.items;
    const {colItems} = lineItems[lineIndex];
    let x = colItems[colIndex].left - this.contentEle.scrollLeft;
    let y = lineIndex * 20 - this.contentEle.scrollTop;
    if (scrollPage) {
        const w = this.pageEle.w;
        const h = this.pageEle.h;
        if (y + 20 > h) {
            const delta = y + 20 - h;
            const top = this.contentEle.scrollTop + delta;
            this.contentEle.scrollTop = top;
            y = lineIndex * 20 - top;
        } else if (y < 0) {
            const top = this.contentEle.scrollTop + y - 0;
            this.contentEle.scrollTop = top;
            y = lineIndex * 20 - top;
        }
        if (x + 5 > w) {
            const delta = x + 5 - w;
            const left = this.contentEle.scrollLeft + delta;
            this.contentEle.scrollLeft = left;
            x = colItems[colIndex].left - left;
        } else if (x < 0) {
            const left = this.contentEle.scrollLeft + x - 0;
            this.contentEle.scrollLeft = left;
            x = colItems[colIndex].left - left;
        }
        this.fakeInputEle.ref.focus();
    }
    this.cursorEle.x = x;
    this.cursorEle.y = y;
    this.activeLineEle.y = y;
    this.activeLineEle.v = 1;

    if (this.hasVisibleRange()) {
        this.refreshHighlight([]);
    } else {
        this.highlightWordsByCursor();
    }

    webapp.util.startAnimation(this.cursorEle);
}
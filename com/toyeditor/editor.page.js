function createPageContent(text) {
    const {page} = this.data;
    // lines
    page.lines = text.split('\n').map(line => Array.from(line));
    // items
    page.items = [];
    for (let i = 0; i < page.lines.length; i++) {
        this._insertPageItemAfter(i - 1);
    }

    // undo
    const lineIndex = page.lines.length - 1;
    const colIndex = page.lines[lineIndex].length - 1;
    const trans = this._pushUndo(
        webapp.constant.createPageCommand,
        text,
        this.newRange(
            this.newPosition(0, 0),
            this.newPosition(lineIndex, colIndex),
        ),
    );
    // publish
    this.publishChange(trans);
    // set
    this._updatePage();

    // test
    this.pageTokenize();
}

function pageInsertText(text) {
    const {page} = this.data;
    const range = this.getRange();
    const {lineIndex, colIndex} = range.start;
    const lines = text.split('\n').map(line => Array.from(line));
    const lineChars = page.lines[lineIndex];
    const newCursorPos = this.newPosition(lineIndex + lines.length - 1, 0);
    // lines
    for (let i = 0; i < lines.length; i++) {
        let chars = [];
        if (i === 0) {
            chars = lineChars.slice(0, colIndex);
        }
        chars = chars.concat(lines[i]);
        if (i === lines.length - 1) {
            newCursorPos.colIndex = chars.length;
            const tmp = lineChars.slice(colIndex);
            chars = chars.concat(tmp);
        }
        if (i === 0) {
            page.lines[lineIndex] = chars;
        } else {
            page.lines.splice(lineIndex + i, 0, chars);
        }
    }
    // items
    if (lines.length === 1) {
        this._updatePageItem(lineIndex);
    } else {
        for (let i = 0; i < lines.length; i++) {
            if (i !== 0) {
                page.items.splice(lineIndex + i, 0, {});
            }
            this._updatePageItem(lineIndex + i);
        }
    }
    // undo
    const trans = this._pushUndo(
        webapp.constant.insertTextCommand,
        text,
        this.newRange(range.start, newCursorPos),
    );
    // publish
    this.publishChange(trans);
    // set
    this._updatePage();

    // test
    this.pageTokenize();

    return newCursorPos;
}

function pageGetText() {
    const {page} = this.data;
    const {start, end} = this.getRange();
    // 没有拖蓝的情况
    if (this.samePosition(start, end)) {
        return page.lines[start.lineIndex].join('');
    } else if (start.lineIndex === end.lineIndex) {
        return page.lines[start.lineIndex].slice(start.colIndex, end.colIndex).join('');
    } else {
        let selectedLines = [page.lines[start.lineIndex].slice(start.colIndex).join('')];
        for (let i = start.lineIndex + 1; i < end.lineIndex; i++) {
            selectedLines.push(page.lines[i].join(''));
        }
        selectedLines.push(page.lines[end.lineIndex].slice(0, end.colIndex).join(''));
        return selectedLines.join('\n');
    }
}

function pageDeleteText() {
    const {page} = this.data;
    const {start, end} = this.getRange();
    let deletedText = '';
    let newCursorPos = this.copyPosition(start);
    // 没有拖蓝的情况
    if (this.samePosition(start, end)) {
        const {lineIndex, colIndex} = start;
        const lineChars = page.lines[lineIndex];
        if (colIndex > 0) {
            // lines
            const c = lineChars.splice(colIndex - 1, 1);
            deletedText = c[0];
            // items
            this._updatePageItem(lineIndex);
            // cursor
            newCursorPos.colIndex = colIndex - 1;
        } else if (lineIndex > 0) {
            // lines
            const prevLineChars = page.lines[lineIndex - 1];
            // cursor
            newCursorPos.lineIndex = lineIndex - 1;
            newCursorPos.colIndex = prevLineChars.length;
            page.lines[lineIndex - 1] = prevLineChars.concat(lineChars);
            page.lines.splice(lineIndex, 1);
            deletedText = '\n';
            // items
            this._deletePageItemAfter(lineIndex - 1);
            this._updatePageItem(lineIndex - 1);
        }
    } else {
        // 有拖蓝的情况，从后往前更新
        if (start.lineIndex === end.lineIndex) {
            deletedText = page.lines[start.lineIndex].slice(start.colIndex, end.colIndex).join('');
        } else {
            let deletedLines = [page.lines[start.lineIndex].slice(start.colIndex).join('')];
            for (let i = start.lineIndex + 1; i < end.lineIndex; i++) {
                deletedLines.push(page.lines[i].join(''));
            }
            deletedLines.push(page.lines[end.lineIndex].slice(0, end.colIndex).join(''));
            deletedText = deletedLines.join('\n');
        }
        let prevChars = page.lines[start.lineIndex].slice(0, start.colIndex);
        let nextChars = page.lines[end.lineIndex].slice(end.colIndex);
        const newLineChars = prevChars.concat(nextChars);
        // lines
        page.lines.splice(
            start.lineIndex,
            end.lineIndex - start.lineIndex + 1,
            newLineChars,
        );
        // items
        page.items.splice(
            start.lineIndex,
            end.lineIndex - start.lineIndex,
        );
        this._updatePageItem(start.lineIndex);
    }
    // undo
    const trans = this._pushUndo(
        webapp.constant.deleteTextCommand,
        deletedText,
        this.newRange(newCursorPos, end),
    );
    // publish
    this.publishChange(trans);
    // set
    this._updatePage();

    // test
    this.pageTokenize();

    return {
        newCursorPos: newCursorPos,
        deletedText: deletedText,
    };
}

function pageTokenize() {
    const {page} = this.data;
    const {items} = page;
    const code = items.map(i => i.chars.join('')).join('\n');
    if (!code) {
        return;
    }
    webapp.worker.call('tokenize', ['js', code])
        .then(data => {
            const newCode = items.map(i => i.chars.join('')).join('\n');
            if (code !== newCode) {
                return;
            }
            const tokens = data.tokens;
            page.brackets = data.brackets;
            for (let i=0; i<items.length; i++) {
                const item = items[i];
                if (i < tokens.length) {
                    item.tokenItems = [];
                    const tokensOfLine = tokens[i];
                    item.key = `(tokenized)${item.chars.join('')}`;
                    let index = 0;
                    for (let j = 0; j < tokensOfLine.length; j++) {
                        const token = tokensOfLine[j];
                        if (index < token.start) {
                            const text = item.chars.slice(index, token.start).join('');
                            item.tokenItems.push({
                                key: `${index}text`,
                                text: text,
                                start: index,
                                end: token.start,
                                kind: 'default',
                                color: webapp.theme.tokenColor.default,
                            });
                        }
                        const text = item.chars.slice(token.start, token.end).join('');
                        item.tokenItems.push({
                            key: `${token.start}text`,
                            text: text,
                            start: token.start,
                            end: token.end,
                            kind: token.kind,
                            color: webapp.theme.tokenColor[token.kind],
                        });
                        index = token.end;
                    }
                    if (index < item.chars.length) {
                        const text = item.chars.slice(index).join('');
                        item.tokenItems.push({
                            key: `${index}text`,
                            text: text,
                            start: index,
                            end: item.chars.length,
                            kind: 'default',
                            color: webapp.theme.tokenColor.default,
                        });
                    }
                }
                items[i] = Object.assign({}, item);
            }
            page.items = [...page.items];
            this._updatePage();
        })
        .catch(e => {
            webapp.log.debug(e);
        });
}

function _deletePageItemAfter(index) {
    this.data.page.items.splice(index + 1, 1);
}

function _insertPageItemAfter(index) {
    this.data.page.items.splice(index + 1, 0, {});
    this._updatePageItem(index + 1);
}

function _updatePageItem(index) {
    const {page} = this.data;
    const chars = page.lines[index];
    const lineNo = index + 1 + '';
    const ret = {
        key: '',
        chars: chars.map(char => webapp.font.isSupportedMonoChar(char) ? char : '\ufffd'),
        lineNo: lineNo,
        lineNoWidth: lineNo.length * 36 / 5,
        colItems: [],
        wordItems: [],
        tokenItems: [],
        width: 0,
    };
    // make cols
    let lastX2 = 0;
    for (let i = 0; i <= chars.length; i++) {
        let w = 5;
        if (i !== chars.length) {
            w = webapp.util.monoTextWidth(chars[i], this.data.fontSize);
        }
        const colItem = {
            left: lastX2,
            width: w,
            right: lastX2 + w,
        };
        lastX2 = colItem.right;
        ret.colItems.push(colItem);
    }
    // make spans
    const text = ret.chars.join('');
    ret.tokenItems = [{key: text, text: text, start: 0, end: ret.chars.length}];
    ret.width = lastX2;
    page.items[index] = ret;
    page.items[index].key = text;
}

function _updatePage() {
    const {items, highlightItems} = this.data.page;
    this.data.lineNoMaxWidth = (items.length + '').length;
    this.leftEle.items = items;
    this.contentEle.items = items;
    this.pageLineCount = items.length;

    this.highlightWordsByCursor();
}

function refreshPage() {
    this.leftEle._updateList();
    this.contentEle._updateList();
}
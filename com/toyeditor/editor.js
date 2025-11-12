function _() {
    this.activeLineEle = undefined;
    this.leftEle = undefined;
    this.pageEle = undefined;
    this.selectionEle = undefined;
    this.highlightEle = undefined;
    this.contentEle = undefined;
    this.cursorEle = undefined;
    this.fakeInputEle = undefined;
}

function onCreated() {
    this.data.compositionStartCursor = this.newPosition(-1, -1);
    this.data.fontSize = 12;
    this.setValue('');
}

function computeLineNo(container, index) {
    const lineNo = index + 1 + '';
    const {fontSize} = container.root.data;
    return {
        index,
        key: lineNo,
        lineNo: lineNo,
        width: lineNo.length * 0.6 * fontSize,
        x: 0,
        y: index * 20,
        w: container.w,
        h: 20,
    };
}

function computeLine(container, index) {
    const data = container.items[index];
    return Object.assign(data, {
        index,
        key: data.key,
        x: 0,
        y: index * 20,
        w: data.width,
        h: 20,
    });
}

function computeSpan(container, index) {
    const data = container.items[index];
    return Object.assign(data, {
        index,
        key: data.key,
        x: 0,
        y: 0,
        w: 0,
        h: container.h,
    });
}

function computeSelectionItem(container, index) {
    const selectionItem = container.items[index];
    const {lineIndex, start, end} = selectionItem;
    const ret = {index, key: `${lineIndex}:${start}-${end}`, x: 0, y: lineIndex * 20, w: 0, h: 20};
    const {items} = container.root.data.page;
    if (lineIndex < items.length) {
        const {colItems} = items[lineIndex];
        ret.x = colItems.length > start ? colItems[start].left : 0;
        const end2 = end === -1 ? colItems.length - 1 : Math.min(end, colItems.length - 1);
        ret.w = colItems.length > end2 ? colItems[end2].left - ret.x : 0;
    }
    return Object.assign(selectionItem, ret);
}

function setValue(value) {
    this.doCreatePage(value);
}

function getValue() {
    const {page} = this.data;
    return page.lines.map(chars => chars.join('')).join('\n');
}

function onValueChanged(fun) {
    this.data.listeners.push(fun);
}

function publishChange(trans) {
    this.data.listeners.forEach(listener => listener(trans));
}

function handlePageContextMenu(ele, ev) {
    ev.preventDefault();
    const menu = [
        'Show Context Actions',
        '',
        'Paste',
        'Copy / Paste Special',
        'Column Selection Mode',
        '',
        'Go To',
        '',
        'Folding',
        '',
        'Edit as Table...',
        '',
        'Refactor',
        'Generate...',
        '',
        'Open In',
        '',
        'Local History',
        'Git',
        '',
        'Go Tools',
        'Compare with Clipboard',
        'Diagrams',
    ];
    webapp.menu?.show(menu, ev);
}
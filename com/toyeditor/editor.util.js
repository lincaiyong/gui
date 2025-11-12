function newPosition(line, col) {
    return {lineIndex: line, colIndex: col};
}

function copyPosition(pos) {
    return {
        lineIndex: pos.lineIndex,
        colIndex: pos.colIndex,
    };
}

function samePosition(a, b) {
    return a.lineIndex === b.lineIndex && a.colIndex === b.colIndex;
}

function newRange(start, end) {
    return {
        start: this.copyPosition(start),
        end: this.copyPosition(end),
    };
}

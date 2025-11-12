function refreshHighlight(highlightItems) {
    const items = [];
    highlightItems.forEach(({range, color = webapp.theme.editorHighlightColor}) => {
        const {start, end} = range;
        if (start.lineIndex === end.lineIndex) {
            items.push({lineIndex: start.lineIndex, start: start.colIndex, end: end.colIndex, color});
        } else {
            items.push({lineIndex: start.lineIndex, start: start.colIndex, end: -1, color});
            for (let i = start.lineIndex + 1; i < end.lineIndex; i++) {
                items.push({lineIndex: i, start: 0, end: -1, color});
            }
            items.push({lineIndex: end.lineIndex, start: 0, end: end.colIndex, color});
        }
    });
    this.highlightEle.items = items;
    this.data.page.highlightItems = items;
}

function highlightWordsByCursor() {
    this.refreshHighlight([]);

    let result = [];
    let {left, right} = this.getWordByCursor();
    if (left?.isWordLike || right?.isWordLike) {
        const word = left?.isWordLike ? left : right;
        const {page} = this.data;
        const {items} = page;

        for (let i = 0; i < items.length; i++) {
            const tokens = items[i].tokenItems;
            for (let j = 0; j < tokens.length; j++) {
                const token = tokens[j];
                if (token.text === word.value) {
                    result.push({
                        range: this.newRange(this.newPosition(i, token.start), this.newPosition(i, token.end)),
                    });
                }
            }
        }
    }
    if (result.length === 1) {
        result = [];
    }

    const pairResult = [];
    const pairs = {'(': ')', '[': ']', '{': '}'};
    if (pairs[left?.value] || pairs[right?.value]) {
        const brackets = this.data.page.brackets;
        let {lineIndex, colIndex} = this.data.cursor.position;
        let openBracket;
        if (pairs[right?.value]) {
            openBracket = right.value;
        } else {
            openBracket = left.value;
            colIndex--;
        }
        const closeBracket = pairs[openBracket];

        let leftPair, rightPair;
        let depth = 1;
        for (let i = 0; i < brackets.length; i++) {
            const bracket = brackets[i];
            if (leftPair) {
                if (bracket.kind === openBracket) {
                    depth++;
                } else if (bracket.kind === closeBracket) {
                    depth--;
                    if (depth === 0) {
                        rightPair = bracket;
                        break;
                    }
                }
            } else if (bracket.line === lineIndex && bracket.start === colIndex) {
                leftPair = bracket;
            }
        }
        if (leftPair && rightPair) {
            pairResult.push({
                range: this.newRange(this.newPosition(leftPair.line, leftPair.start), this.newPosition(leftPair.line, leftPair.end)),
                color: webapp.theme.editorBracketHighlightColor,
            });
            pairResult.push({
                range: this.newRange(this.newPosition(rightPair.line, rightPair.start), this.newPosition(rightPair.line, rightPair.end)),
                color: webapp.theme.editorBracketHighlightColor,
            });
        }
    }

    if (pairResult.length === 0) {
        const pairs2 = {')': '(', ']': '[', '}': '{'};
        if (pairs2[left?.value] || pairs2[right?.value]) {
            const brackets = this.data.page.brackets;
            let {lineIndex, colIndex} = this.data.cursor.position;
            let openBracket;
            if (pairs2[right?.value]) {
                openBracket = right.value;
            } else {
                openBracket = left.value;
                colIndex--;
            }
            const closeBracket = pairs2[openBracket];

            let leftPair, rightPair;
            let depth = 1;
            for (let i = brackets.length - 1; i >= 0; i--) {
                const bracket = brackets[i];
                if (leftPair) {
                    if (bracket.kind === openBracket) {
                        depth++;
                    } else if (bracket.kind === closeBracket) {
                        depth--;
                        if (depth === 0) {
                            rightPair = bracket;
                            break;
                        }
                    }
                } else if (bracket.line === lineIndex && bracket.start === colIndex) {
                    leftPair = bracket;
                }
            }
            if (leftPair && rightPair) {
                result.push({
                    range: this.newRange(this.newPosition(leftPair.line, leftPair.start), this.newPosition(leftPair.line, leftPair.end)),
                    color: webapp.theme.editorBracketHighlightColor,
                });
                result.push({
                    range: this.newRange(this.newPosition(rightPair.line, rightPair.start), this.newPosition(rightPair.line, rightPair.end)),
                    color: webapp.theme.editorBracketHighlightColor,
                });
            }
        }
    } else {
        result = result.concat(pairResult);
    }

    if (result.length > 1) {
        this.refreshHighlight(result);
    }
}
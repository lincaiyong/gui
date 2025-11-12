function handleKeyUp(ele, ev) {
    // console.debug('keyup', ev);
    const {keyPressed} = ele.root.data;
    if (ev.code === 'ShiftLeft' || ev.code === 'ShiftRight') {
        keyPressed.shift = false;
    }
}

function handleKeyDown(ele, ev) {
    // console.debug('keydown', ev);
    // control: ev.ctrlKey,
    // shift: ev.shiftKey,
    // option: ev.altKey,
    // command: ev.metaKey,
    const {metaKey:command, ctrlKey:control, altKey:option, shiftKey:shift} = ev;
    const {anchor, keyPressed, page} = this.data;
    if (ev.code === 'ArrowLeft' || ev.code === 'ArrowRight' || ev.code === 'ArrowUp' || ev.code === 'ArrowDown') {
        if (shift) {
            if (!anchor.valid) {
                this.createAnchor();
            }
        }
        if (ev.code === 'ArrowLeft') {
            if (command) {
                this.moveCursorLeftByLine();
            } else if (option) {
                this.moveCursorLeftByWord();
            } else if (shift) {
                this.moveCursorLeftIgnoreRange();
            } else {
                this.moveCursorLeft();
            }
        } else if (ev.code === 'ArrowRight') {
            if (command) {
                this.moveCursorRightByLine();
            } else if (option) {
                this.moveCursorRightByWord();
            } else if (shift) {
                this.moveCursorRightIgnoreRange();
            } else {
                this.moveCursorRight();
            }
        } else if (ev.code === 'ArrowUp') {
            this.moveCursorUp();
        } else {
            this.moveCursorDown();
        }
        if (!shift) {
            this.clearAnchor();
        }
        this.refreshCursor(true);
        this.updateSelection();
        this.refreshSelection();
    } else if (ev.code === 'Backspace') {
        if (command) {
            const {start, end} = this.getRange();
            if (start.lineIndex > 0) {
                start.lineIndex = start.lineIndex - 1;
                start.colIndex = page.lines[start.lineIndex].length;
            } else {
                start.colIndex = 0;
            }
            this.moveCursorToPosition(start);
            this.createAnchor();
            end.colIndex = page.lines[end.lineIndex].length;
            this.moveCursorToPosition(end);
            this.doDeleteText();
        } else if (option && !this.hasVisibleRange()) {
            this.createAnchor();
            this.moveCursorLeftByWord();
            this.doDeleteText();
        } else {
            this.doDeleteText();
        }
    } else if (ev.code === 'Enter') {
        this.doInsertText('\n');
    } else if (ev.code === 'KeyA' && command) {
        this.moveCursorToPageStart();
        this.createAnchor();
        this.moveCursorToPageEnd();
        this.refreshCursor(true);
        this.updateSelection();
        this.refreshSelection();
    } else if (ev.code === 'ShiftLeft' || ev.code === 'ShiftRight') {
        keyPressed.shift = true;
    } else if (ev.code === 'KeyZ' && command && shift) {
        this.executeRedo();
    } else if (ev.code === 'KeyZ' && command) {
        this.executeUndo();
    }
}

function handleCompositionUpdate(ele, ev) {
    if (ev.data) {
        const {cursor, compositionStartCursor} = this.data;
        if (compositionStartCursor.lineIndex === -1) {
            this.data.compositionStartCursor = this.copyPosition(cursor.position);
            this.freezeUndo(() => this.doInsertText(ev.data));
        } else {
            this.createAnchor();
            this.moveCursorToPosition(compositionStartCursor);
            this.freezeUndo(() => {
                this.doDeleteText();
                this.doInsertText(ev.data);
            });
        }
    }
}

function handleCompositionEnd(ele, ev) {
    if (ev.data) {
        const {compositionStartCursor} = this.data;
        if (compositionStartCursor.lineIndex !== -1) {
            this.createAnchor();
            this.moveCursorToPosition(compositionStartCursor);
            this.freezeUndo(() => this.doDeleteText());
        }
        this.doInsertText(ev.data);
        compositionStartCursor.lineIndex = -1;
        this.fakeInputEle.ref.value = '';
    }
}

function handleInput(ele, ev) {
    if (ev.data) {
        // console.debug('input', ev);
        if (ev.inputType === 'insertText') {
            this.doInsertText(ev.data);
            this.fakeInputEle.ref.value = '';
        }
    }
}

function handleCopy(ele, ev) {
    ev.clipboardData.setData('text/plain', this.doCopyText());
    ev.preventDefault();
}

function handleCut(ele, ev) {
    ev.clipboardData.setData('text/plain', this.doCutText());
    ev.preventDefault();
}

function handlePaste(ele, ev) {
    this.doPasteText(ev.clipboardData.getData('text/plain'));
    ev.preventDefault();
}

function handleFocus() {
    return () => this.hideCursor();
}

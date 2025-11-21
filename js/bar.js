function bar_handleMouseDown(ele, mouseDownEvent) {
    const prev = ele.prev;
    const next = ele.next;
    const state = {prevX: mouseDownEvent.clientX, prevY: mouseDownEvent.clientY};
    const cancelMouseMoveListener = g.addListener(window, 'mousemove', ev => {
        const safeDist = 10;
        if (ele.cursor === 'col-resize') {
            const newX = ele.x + ev.clientX - state.prevX;
            state.prevX = ev.clientX;
            if (newX < prev.x + safeDist) {
                ele.x = prev.x + safeDist;
            } else if (newX + ele.w > next.x + next.w - safeDist) {
                ele.x = next.x + next.w - safeDist - ele.w;
            } else {
                ele.x = newX;
            }
        } else {
            const newY = ele.y + ev.clientY - state.prevY;
            state.prevY = ev.clientY;
            if (newY < prev.y + safeDist) {
                ele.y = prev.y + safeDist;
            } else if (newY + ele.h > next.y + next.h - safeDist) {
                ele.y = next.y + next.h - safeDist - ele.h;
            } else {
                ele.y = newY;
            }
        }
    });
    g.onceListener(window, 'mouseup', () => {
        cancelMouseMoveListener();
    });
}
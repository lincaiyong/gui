function handleMouseDown(ele, mouseDownEvent) {
    const prev = ele._('prev');
    const next = ele._('next');
    const state = {prevX: mouseDownEvent.clientX, prevY: mouseDownEvent.clientY};
    const cancelMouseMoveListener = g.event.addListener(window, 'mousemove', ev => {
        const safeDist = 10;
        if (ele.cursor === 'col-resize') {
            const newX = ele.x + ev.clientX - state.prevX;
            state.prevX = ev.clientX;
            if (newX < prev.x + safeDist) {
                ele.x = prev.x + safeDist;
            } else if (newX > next.x + next.w - safeDist) {
                ele.x = next.x + next.w - safeDist;
            } else {
                ele.x = newX;
            }
        } else {
            const newY = ele.y + ev.clientY - state.prevY;
            state.prevY = ev.clientY;
            if (newY < prev.y + safeDist) {
                ele.y = prev.y + safeDist;
            } else if (newY > next.y + next.h - safeDist) {
                ele.y = next.y + next.h - safeDist;
            } else {
                ele.y = newY;
            }
        }
    });
    g.event.onceListener(window, 'mouseup', () => {
        cancelMouseMoveListener();
    });
}
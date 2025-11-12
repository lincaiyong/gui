class ScrollbarWrapper {
    constructor(container, vh) {
        this.vertical = vh === 'v';
        this.container = container;
        this.bar = this.vertical ? container.vBarEle : container.hBarEle;
        if (this.vertical) {
            this.bar._properties.y.reset();
            this.bar._properties.h.reset();
        } else {
            this.bar._properties.x.reset();
            this.bar._properties.w.reset();
        }

        const _fade = () => this.active = false;
        this.debouncedFade = page.util.debounce(_fade, container.scrollBarFadeTime);

        page.util.assert(container instanceof Container);
        page.util.assert(this.bar instanceof Scrollbar);
    }
    get contentLen() {
        return this.vertical ? this.container.childHeight : this.container.childWidth;
    }
    get containerLen() {
        return this.vertical ? this.container.h : this.container.w;
    }
    get barPos() {
        return this.vertical ? this.bar.y : this.bar.x;
    }
    set barPos(v) {
        this.vertical ? this.bar.y = v : this.bar.x = v;
    }
    get barLen() {
        return this.vertical ? this.bar.h : this.bar.w;
    }
    set barLen(v) {
        this.vertical ? this.bar.h = v : this.bar.w = v;
    }
    get scrollVal() {
        return this.vertical ? this.container.scrollTop : this.container.scrollLeft;
    }
    set scrollVal(v) {
        this.vertical ? this.container.scrollTop = v : this.container.scrollLeft = v;
    }
    set active(v) {
        if (v) {
            this.bar.opacity = 1;
        } else {
            this.bar.opacity = 0;
        }
    }
    get containerRect() {
        const rect = this.container.ref.getBoundingClientRect();
        const min = this.vertical ? rect.top : rect.left;
        const max = this.vertical ? rect.bottom : rect.right;
        return {min, max};
    }
    show(flag) {
        if (flag && this.contentLen > this.containerLen) {
            const { scrollBarMargin, scrollBarMinLength } = this.container;
            this.barLen = Math.max(this.containerLen ** 2 / this.contentLen, scrollBarMinLength);
            this.barPos = (this.containerLen - this.barLen - 2 * scrollBarMargin) * this.scrollVal / (this.contentLen - this.containerLen) + scrollBarMargin;
            this.bar.v = 1;
        } else {
            this.bar.v = 0;
        }
    }
    getEventPos(ev) {
        return this.vertical ? ev.clientY : ev.clientX;
    }
    initDraggable() {
        this.bar.onMouseDown = (_, ev0) => {
            ev0.stopPropagation();
            const prev = {val: this.getEventPos(ev0)};
            const cancelMouseMove = page.event.addListener(window, 'mousemove', ev => {
                const evPos = this.getEventPos(ev);
                const {min, max} = this.containerRect;
                if ((evPos < min && prev.val === min) || (evPos > max && prev.val === max)) {
                    return;
                }
                const mouse = Math.min(max, Math.max(evPos, min))
                this.scrollVal = Math.min(this.contentLen - this.containerLen, Math.max(this.scrollVal + (mouse - prev.val) / (this.containerLen / this.contentLen), 0));
                prev.val = mouse;
                this.show(true);
                this.active = true;
            });
            page.event.onceListener(window, 'mouseup', () => {
                this.active = false;
                cancelMouseMove();
            });
        };
    }
    handleWheel(ev) {
        if (this.contentLen > this.containerLen) {
            this.active = true;
            const delta = this.vertical ? ev.deltaY : ev.deltaX;
            this.scrollVal= Math.min(this.contentLen - this.containerLen, Math.max(this.scrollVal + delta, 0));
            this.show(true);
            this.debouncedFade();
        }
    }
}
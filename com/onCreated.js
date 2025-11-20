function onCreated() {
    if (!this.list) {
        const child = g.createElement(this.model.slot[0], this);
        this.childWidth = child.w;
        this.childHeight = child.h;
        child.onUpdated((k, v) => {
            if (k === 'w') {
                this.childWidth = v;
            } else if (k === 'h') {
                this.childHeight = v;
            }
        });
    }

    if (this.scrollable) {
        this.hBar = new ScrollbarWrapper(this, 'h');
        this.vBar = new ScrollbarWrapper(this, 'v');
        const bars = [this.hBar, this.vBar];
        bars.forEach(bar => bar.initDraggable());
        this.onWheel = (_, ev) => {
            ev.preventDefault();
            bars.forEach(bar => bar.handleWheel(ev));
        };
    }
}
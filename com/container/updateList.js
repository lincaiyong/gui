function _updateList() {
    if (!this.list) {
        return;
    }
    const data = this.items;

    page.util.assert(data instanceof Array);
    const computeFunc = this.model.slot[0].properties.compute[0]();
    page.util.assert(computeFunc instanceof Function);

    const scrollLeft = this.scrollLeft || 0;
    const scrollTop = this.scrollTop || 0;
    const RESERVED_COUNT = 2;
    let mw = 0;
    let mh = 0;
    const computedItems = [];
    const visible = [];
    let prevItem = null;
    for (let i = 0; i < data.length; i++) {
        const item = computeFunc(this, i, prevItem);
        page.util.assert(typeof(item.key) === 'string');
        computedItems.push(item);
        prevItem = item;

        mw = Math.max(item.x + item.w, mw);
        mh = Math.max(item.y + item.h, mh);

        if (!this.virtual) {
            visible.push(i);
        } else {
            const x = prevItem.x - scrollLeft;
            const x2 = x + prevItem.w;
            const y = prevItem.y - scrollTop;
            const y2 = y + prevItem.h;
            if (!(x > this.w || x2 < 0 || y > this.h || y2 < 0)) {
                visible.push(i);
            }
        }
    }

    if (this.reuseItem) {
        const old = {};
        for (let i = RESERVED_COUNT; i < this.children.length; i++) {
            const child = this.children[i];
            const key = child.data.key;
            page.util.assert(typeof(key) === 'string');
            if (key in old) {
                old[key].push(child);
            } else {
                old[key] = [child];
            }
        }

        const hitKey = {};
        visible.forEach(i => {
            const key = computedItems[i].key;
            page.util.assert(typeof(key) === 'string');
            if (key in old && old[key].length > 0) {
                hitKey[i] = old[key].shift();
            }
        });
        let other = [];
        Object.values(old).forEach(t => other = other.concat(t));

        const nonHitKey = [];
        visible.forEach(i => {
            let child = hitKey[i];
            if (!child) {
                child = other.shift();
                if (!child) {
                    child = page.createElement(this.model.slot[0], this);
                    ['x', 'y', 'w', 'h'].forEach(k => child._properties[k].reset());
                }
                nonHitKey.push(child);
            }
        });
        other.forEach(t => page.removeElement(t));

        page.log.trace(`total: ${visible.length}, hit: ${Object.values(hitKey).length}, non hit: ${nonHitKey.length}`);

        visible.forEach(i => {
            const item = computedItems[i];
            const child = hitKey[i] || nonHitKey.shift();
            child.data = item;
            child.x = item.x - scrollLeft;
            child.y = item.y - scrollTop;
            child.w = item.w;
            child.h = item.h;
        });
    } else {
        while (this.children.length > visible.length + 2) {
            const child = this.children[this.children.length - 1];
            page.removeElement(child);
        }
        while (this.children.length < visible.length + 2) {
            page.createElement(this.model.slot[0], this);
        }
        for (let i = 0; i < visible.length; i++) {
            const child = this.children[i+RESERVED_COUNT];
            const item = computedItems[visible[i]];
            child.data = item;
            child.x = item.x - scrollLeft;
            child.y = item.y - scrollTop;
            child.w = item.w;
            child.h = item.h;
        }
    }

    this.childWidth = this.minWidth > 0 ? Math.max(mw, this.minWidth) : mw;
    this.childHeight = mh;
    if (this.align !== 'none') {
        const w = this.align === 'max' ? this.childWidth : Math.max(this.childWidth, this.cw);
        for (let i = RESERVED_COUNT; i < this.children.length; i++) {
            const child = this.children[i];
            child.w = w;
        }
    }

    if (this.scrollable) {
        if (mw - scrollLeft < this.cw) {
            this.scrollLeft = Math.max(mw - this.cw, 0);
        }
        if (mh - scrollTop < this.ch) {
            this.scrollTop = Math.max(mh - this.ch, 0);
        }
        this.hBar?.show(true);
        this.vBar?.show(true);
    }
}
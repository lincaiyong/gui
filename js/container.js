function container_updateList(ele) {
    if (!ele.list) {
        return;
    }
    const data = ele.items;
    const scrollLeft = ele.scrollLeft || 0;
    const scrollTop = ele.scrollTop || 0;
    const RESERVED_COUNT = 2;
    let mw = 0;
    let mh = 0;
    const computedItems = [];
    const visible = [];
    let prevItem = null;
    for (let i = 0; i < data.length; i++) {
        const item = ele.computeItemFn(ele, i, prevItem);
        computedItems.push(item);
        prevItem = item;

        mw = Math.max(item.x + item.w, mw);
        mh = Math.max(item.y + item.h, mh);

        if (!ele.virtual) {
            visible.push(i);
        } else {
            const x = prevItem.x - scrollLeft;
            const x2 = x + prevItem.w;
            const y = prevItem.y - scrollTop;
            const y2 = y + prevItem.h;
            if (!(x > ele.w || x2 < 0 || y > ele.h || y2 < 0)) {
                visible.push(i);
            }
        }
    }

    if (ele.reuseItem) {
        const old = {};
        for (let i = RESERVED_COUNT; i < ele.children.length; i++) {
            const child = ele.children[i];
            const key = child.data.key;
            if (key in old) {
                old[key].push(child);
            } else {
                old[key] = [child];
            }
        }

        const hitKey = {};
        visible.forEach(i => {
            const key = computedItems[i].key;
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
                    child = g.createElement(null, ele, ele.model.itemModel);
                    ['x', 'y', 'w', 'h'].forEach(k => child.properties[k].reset());
                }
                nonHitKey.push(child);
            }
        });
        other.forEach(t => g.destroyElement(t));
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
        while (ele.children.length > visible.length + 2) {
            const child = ele.children[ele.children.length - 1];
            g.destroyElement(child);
        }
        while (ele.children.length < visible.length + 2) {
            g.createElement(null, ele, ele.model.itemModel);
        }
        for (let i = 0; i < visible.length; i++) {
            const child = ele.children[i + RESERVED_COUNT];
            const item = computedItems[visible[i]];
            child.data = item;
            child.x = item.x - scrollLeft;
            child.y = item.y - scrollTop;
            child.w = item.w;
            child.h = item.h;
        }
    }

    ele.childWidth = ele.minWidth > 0 ? Math.max(mw, ele.minWidth) : mw;
    ele.childHeight = mh;
    if (ele.align !== 'none') {
        const w = ele.align === 'max' ? ele.childWidth : Math.max(ele.childWidth, ele.cw);
        for (let i = RESERVED_COUNT; i < ele.children.length; i++) {
            const child = ele.children[i];
            child.w = w;
        }
    }

    if (ele.scrollable) {
        if (mw - scrollLeft < ele.cw) {
            ele.scrollLeft = Math.max(mw - ele.cw, 0);
        }
        if (mh - scrollTop < ele.ch) {
            ele.scrollTop = Math.max(mh - ele.ch, 0);
        }
        ele.hBar.show(true);
        ele.vBar.show(true);
    }
}

function container_handleUpdated(ele, k) {
    // items
    if (k === 'items' && ele.list) {
        container_updateList(ele);
    }

    // scroll
    if (ele.list && ele.virtual && ele.items instanceof Array) {
        if ((k === 'scrollLeft' || k === 'scrollTop') && ele.items instanceof Array) {
            container_updateList(ele);
        }
    } else if (ele.list) {
        const RESERVED_COUNT = 2;
        if (k === 'scrollLeft') {
            for (let i = RESERVED_COUNT; i < ele.children.length; i++) {
                const child = ele.children[i];
                child.x = child.data.x - ele.scrollLeft;
            }
        } else if (k === 'scrollTop') {
            for (let i = RESERVED_COUNT; i < ele.children.length; i++) {
                const child = ele.children[i];
                child.y = child.data.y - ele.scrollTop;
            }
        }
    }

    // w & h -> 影响scroll
    if (ele.scrollable) {
        if ((k === 'w' || k === 'h') && ele.items instanceof Array) {
            container_updateList(ele);
        }
    }
}

function container_handleCreated(ele) {
    if (!ele.list) {
        const child = g.createElement(null, ele, ele.model.itemModel);
        ele.childWidth = child.w;
        ele.childHeight = child.h;
        child.onUpdated = (k, v) => {
            if (k === 'w') {
                ele.childWidth = v;
            } else if (k === 'h') {
                ele.childHeight = v;
            }
        };
    }

    if (ele.scrollable) {
        ele.hBar = new Scrollbar(ele, 'h');
        ele.vBar = new Scrollbar(ele, 'v');
        const bars = [ele.hBar, ele.vBar];
        bars.forEach(bar => bar.initDraggable());
        ele.onWheel = (_, ev) => {
            ev.preventDefault();
            bars.forEach(bar => bar.handleWheel(ev));
        };
    }
}
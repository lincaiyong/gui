function onUpdated(k) {
    // items
    if (k === 'items' && this.list) {
        this._updateList();
    }

    // scroll
    if (this.list && this.virtual && this.items instanceof Array) {
        if ((k === 'scrollLeft' || k === 'scrollTop') && this.items instanceof Array){
            this._updateList();
        }
    } else if (this.list) {
        const RESERVED_COUNT = 2;
        if (k === 'scrollLeft') {
            for (let i = RESERVED_COUNT; i < this.children.length; i++) {
                const child = this.children[i];
                child.x = child.data.x - this.scrollLeft;
            }
        } else if (k === 'scrollTop') {
            for (let i = RESERVED_COUNT; i < this.children.length; i++) {
                const child = this.children[i];
                child.y = child.data.y - this.scrollTop;
            }
        }
    }

    // w & h -> 影响scroll
    if (this.scrollable) {
        if ((k === 'w' || k === 'h') && this.items instanceof Array) {
            this._updateList();
        }
    }
}
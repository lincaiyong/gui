function onUpdated(k, v) {
    if (k === 'items') {
        this.nodeMap = this.makeNodeMap(v);
        this.containerEle.items = this.nodeToItems(this.nodeMap, '', 0, 0);
        this.selectedEle.v = 0;
    }
}
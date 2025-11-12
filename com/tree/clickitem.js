function clickItem(itemEle, ev) {
    const treeEle = itemEle._('this');
    treeEle.selectChild(itemEle, true);
    // 通知发生点击事件
    if (treeEle.onClickItem instanceof Function) {
        treeEle.onClickItem(itemEle, ev);
    }
    // 目录展开折叠
    if (!itemEle.data.leaf) {
        const {key} = itemEle.data;
        const node = treeEle.nodeMap[key];
        node.collapsed = !node.collapsed;
        treeEle.containerEle.items = treeEle.nodeToItems(treeEle.nodeMap, '', 0, 0);
    }
    // 处理blur
    treeEle.onClickOutside = (_, event) => {
        if (ev !== event) {
            treeEle.focus = false;
            return true;
        }
        return false;
    };
}
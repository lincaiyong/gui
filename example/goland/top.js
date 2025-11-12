function handleClickProjectBtn(ele, ev) {
    const { explorerPane } = ele.root;
    ele.selected = !ele.selected;
    explorerPane.v = explorerPane.v ? 0 : 1;
}
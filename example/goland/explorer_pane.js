function onCreated() {
    webapp.util.fetch('/data?name=files').then(v => {
        const resp = JSON.parse(v);
        this.treeEle.items = resp.data;
    }).catch(e => webapp.log.error(e));
}

function handleClickItem(ele, ev) {
    if (ele.data.leaf) {
        webapp.util.fetch(`/data?name=file:${ele.data.key}`).then(v => {
            const resp = JSON.parse(v);
            this.root.editorEle.setValue(resp.data);
        }).catch(e => webapp.log.error(e));
    }
}

function handleClickOptions(ele, ev) {
    webapp.menu?.show(['foo', 'bar', 'a', '', 'b'], ev, this);
    this.menuOpened = true;
    let cancel;
    const handleMenuHide = () => {
        this.menuOpened = false;
        cancel();
    }
    cancel = webapp.menu?._properties.v.onUpdated(handleMenuHide);
}

function handleHover() {

}
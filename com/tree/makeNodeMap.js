function makeNodeMap(items) {
    if (this.sort) {
        items.sort();
    }
    const nodeMap = {};
    nodeMap[''] = {
        parent: null,
        key: '',
        text: '',
        children: [],
        collapsed: false,
    };
    items.forEach(item => {
        let key = '';
        item.split('/').forEach(tmp => {
            const parent = nodeMap[key];
            key = key ? [key, tmp].join('/') : tmp;
            if (!nodeMap[key]) {
                nodeMap[key] = {
                    parent: parent,
                    key: key,
                    text: tmp,
                    children: [],
                    collapsed: true,
                };
                parent.children.push(nodeMap[key]);
            }
        });
    });
    return nodeMap;
}

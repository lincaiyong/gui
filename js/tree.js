function tree_computeItem(container, index) {
    const data = container.items[index];
    const h = container.root.itemHeight;
    return Object.assign(data, {
        index,
        key: data.key,
        x: 0,
        y: index * h,
        w: data.depth * 20 + g.textWidth(data.text, container.fontFamily, 12) + 40,
        h,
    });
}

function tree_clickItem(itemEle, ev) {
    const treeEle = itemEle.local;
    treeEle.handleChildSelected(itemEle, true);
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

function tree_updateItem(itemEle, k, v) {
    if (k === 'data') {
        if (v.leaf) {
            const ext = v.key.substring(v.key.lastIndexOf('.') + 1);
            let src = 'res/svg/text.svg';
            switch (ext) {
                case 'go':
                    src = 'res/svg/go.svg';
                    break;
                case 'js':
                    src = 'res/svg/js.svg';
                    break;
                case 'py':
                    src = 'res/svg/python.svg';
                    break;
            }
            if (v.key === 'go.mod' || v.key.endsWith('/go.mod')) {
                src = 'res/svg/goMod.svg'
            }
            if (v.key === '.gitignore' || v.key.endsWith('/.gitignore')) {
                src = 'res/svg/ignored.svg'
            }
            itemEle.iconEle.src = src;
        } else {
            itemEle.iconEle.src = 'res/svg/folder.svg';
        }
    }
}

function tree_makeNodeMap(items, sort) {
    if (sort) {
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

function tree_nodeToItems(nodeMap, key, index, depth) {
    const node = nodeMap[key];
    if (!node || !node.children || node.collapsed) {
        return [];
    }

    let ret = [];
    const children = node.children;
    let directories = [];
    let files = [];

    // 先将子节点分为目录和文件两类
    for (let i = 0; i < children.length; i++) {
        const childNode = children[i];
        const isLeaf = childNode.children.length === 0;

        const item = {
            index: null, // 暂时不设置索引，后面再统一设置
            key: childNode.key,
            depth: depth,
            leaf: isLeaf,
            collapsed: childNode.collapsed,
            text: childNode.text,
        };

        if (isLeaf) {
            files.push(item);
        } else {
            directories.push(item);
        }
    }

    // 合并目录和文件，目录在前
    const sortedChildren = [...directories, ...files];

    // 处理每个子节点及其子树
    for (let i = 0; i < sortedChildren.length; i++) {
        const item = sortedChildren[i];
        item.index = index++;
        ret.push(item);

        // 如果不是叶子节点且未折叠，则递归处理其子节点
        if (!item.leaf && !item.collapsed) {
            const tmp = tree_nodeToItems(
                nodeMap,
                item.key,
                index,
                depth + 1,
            );
            ret = ret.concat(tmp);
            index += tmp.length;
        }
    }

    return ret;
}

function tree_sortChildren(node) {
    if (node.children.length > 0) {
        node.children.sort((a, b) => {
            if (!!a.children.length === !!b.children.length) {
                return a.key.localeCompare(b.key);
            }
            return a.children.length > 0 ? -1 : 1;
        });
        node.children.forEach(tree_sortChildren);
    }
}
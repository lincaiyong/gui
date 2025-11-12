function nodeToItems(nodeMap, key, index, depth) {
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
            const tmp = this.nodeToItems(
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
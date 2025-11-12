function sortChildren(node) {
    if (node.children.length > 0) {
        node.children.sort((a, b) => {
            if (!!a.children.length === !!b.children.length) {
                return a.key.localeCompare(b.key);
            }
            return a.children.length > 0 ? -1 : 1;
        });
        node.children.forEach(Tree.sortChildren);
    }
}
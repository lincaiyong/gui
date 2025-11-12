function computeItem(container, index) {
    const data = container.items[index];
    const h = container.root.itemHeight;
    return Object.assign(data, {
        index,
        key: data.key,
        x: 0,
        y: index * h,
        w: data.depth * 20 + page.util.textWidth(data.text, container.fontFamily, 12) + 40,
        h,
    });
}
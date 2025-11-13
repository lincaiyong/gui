function handleActive(ele) {
    if (ele.selected) {
        return;
    }
    const oldBgColor = ele.backgroundColor;
    ele.backgroundColor = g.theme.buttonActiveBgColor;
    return () => {
        ele.backgroundColor = oldBgColor;
    };
}

function handleHover(ele, hover) {
    if (hover) {
        this.backgroundColor = this.selected ? g.theme.buttonSelectedBgColor : g.theme.buttonHoverBgColor;
    } else {
        this.backgroundColor = this.selected ? g.theme.buttonSelectedBgColor : g.theme.buttonBgColor;
    }
}
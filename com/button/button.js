function handleActive(ele) {
    if (ele.selected) {
        return;
    }
    const oldBgColor = ele.backgroundColor;
    ele.backgroundColor = page.theme.buttonActiveBgColor;
    return () => {
        ele.backgroundColor = oldBgColor;
    };
}

function handleHover(ele, hover) {
    if (hover) {
        this.backgroundColor = this.selected ? page.theme.buttonSelectedBgColor : page.theme.buttonHoverBgColor;
    } else {
        this.backgroundColor = this.selected ? page.theme.buttonSelectedBgColor : page.theme.buttonBgColor;
    }
}
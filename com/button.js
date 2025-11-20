function button_handleActive(ele) {
    if (ele.selected) {
        return;
    }
    const oldBgColor = ele.backgroundColor;
    ele.backgroundColor = '#DFE1E4';
    return () => {
        ele.backgroundColor = oldBgColor;
    };
}

function button_handleHover(ele, hover) {
    if (hover) {
        this.backgroundColor = this.selected ? '#3475F0' : '#EBECF0';
    } else {
        this.backgroundColor = this.selected ? '#3475F0' : '';
    }
}
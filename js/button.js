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
        ele.backgroundColor = ele.selected ? '#3475F0' : '#EBECF0';
    } else {
        ele.backgroundColor = ele.selected ? '#3475F0' : '';
    }
}
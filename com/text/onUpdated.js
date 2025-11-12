function onUpdated(k, v) {
    switch (k) {
        case 'text':
            this.ref.innerText = v;
            break;
        case 'align':
            if (v === 'left') {
                // this.ref.style.left = '0px';
            } else if (v === 'right') {

            } else {

            }
            break;
        default:
            break;
    }
}

function createAnchor() {
    const {anchor, cursor} = this.data;
    anchor.position = this.copyPosition(cursor.position);
    anchor.valid = true;
}

function clearAnchor() {
    const {anchor} = this.data;
    anchor.valid = false;
}
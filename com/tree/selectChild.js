function selectChild(child, focus) {
    this.selectedChildTop = child.y + this.containerEle.scrollTop;
    this.selectedEle.v = 1;
    this.focus = focus;
}
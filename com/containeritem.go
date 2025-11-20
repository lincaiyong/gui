package com

func ContainerItem(children ...*Element) *Element {
	ret := NewElement("container_item", "div", children...)
	ret.Y("0").X("0").H("0")
	ret.SetProperty("compute", "null").
		SetProperty("update", "null").
		SetProperty("click", "null").
		SetMethod("onUpdated", `function(k, v) {
    this.update?.(this, k, v);
}`)
	return ret
}

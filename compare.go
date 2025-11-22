package gui

func Compare(opt *Opt) *Element {
	ret := NewElement(ElementTypeCompare, ElementTagDiv)
	opt.OnCreated("compare_handleCreated").OnDestroy("compare_handleDestroy").Init(ret)
	return ret
}

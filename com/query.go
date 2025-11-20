package com

func NewQueryOpt() QueryOpt {
	ret := QueryOpt{}
	return ret
}

type QueryOpt struct {
	tag string
}

func (e *Element) QueryChild(queryOpt QueryOpt) []int {
	return nil
}

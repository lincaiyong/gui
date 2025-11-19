package com

func NewBaseOpt() *BaseOpt {
	return &BaseOpt{
		properties: make(map[string]string),
	}
}

type BaseOpt struct {
	properties map[string]string
}

func (o *BaseOpt) SetProperty(key, value string) {
	o.properties[key] = value
}

func (o *BaseOpt) Properties() map[string]string {
	return o.properties
}

func (o *BaseOpt) Init(e *Element) {
	for k, v := range o.properties {
		e.SetProperty(k, v)
	}
}

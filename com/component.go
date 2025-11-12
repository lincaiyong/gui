package com

type Component interface {
	ExtraInfo() *ExtraInfo
	Tag() string
	Name() string
	Props() map[string]string
	Children() []Component
	Slots() []Component
}

type ExtraInfo struct {
	name          string
	properties    []string
	methods       []string
	staticMethods []string
	bindJs        map[string]string
	defaultValue  map[string]string
	thisComponent Component
	selfIndex     []int
}

func (e *ExtraInfo) Name() string {
	return e.name
}

func (e *ExtraInfo) SetName(name string) {
	e.name = name
}

func (e *ExtraInfo) Properties() []string {
	return e.properties
}

func (e *ExtraInfo) SetProperties(properties []string) {
	e.properties = properties
}

func (e *ExtraInfo) Methods() []string {
	return e.methods
}

func (e *ExtraInfo) AddMethod(method string) {
	e.methods = append(e.methods, method)
}

func (e *ExtraInfo) StaticMethods() []string {
	return e.staticMethods
}

func (e *ExtraInfo) AddStaticMethod(method string) {
	e.staticMethods = append(e.staticMethods, method)
}

func (e *ExtraInfo) GetBindJs(k string) string {
	return e.bindJs[k]
}

func (e *ExtraInfo) SetBindJs(k, v string) {
	e.bindJs[k] = v
}

func (e *ExtraInfo) GetDefaultValue(k string) string {
	return e.defaultValue[k]
}

func (e *ExtraInfo) SetDefaultValue(k, v string) {
	e.defaultValue[k] = v
}

func (e *ExtraInfo) ThisComponent() Component {
	return e.thisComponent
}

func (e *ExtraInfo) SetThisComponent(thisComponent Component) {
	e.thisComponent = thisComponent
}

func (e *ExtraInfo) SelfIndex() []int {
	return e.selfIndex
}

func (e *ExtraInfo) SetSelfIndex(selfIndex []int) {
	e.selfIndex = selfIndex
}

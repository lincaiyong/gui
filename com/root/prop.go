package root

var rootProps = map[string]string{}

func AddProp(k, v string) {
	rootProps[k] = v
}

func Props() map[string]string {
	return rootProps
}

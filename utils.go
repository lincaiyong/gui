package gui

import (
	"reflect"
	"sort"
	"unicode"
)

func pascalCase(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func sortedKeys(m any) []string {
	keys := make([]string, 0)
	for _, k := range reflect.ValueOf(m).MapKeys() {
		keys = append(keys, k.String())
	}
	sort.Strings(keys)
	return keys
}

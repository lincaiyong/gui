package js

import (
	"regexp"
)

var data = map[string]map[string]string{}

func Set(comName string, code ...string) {
	if data[comName] == nil {
		data[comName] = map[string]string{}
	}
	re := regexp.MustCompile(`(?ms)^function\s+([a-zA-Z_][a-zA-Z0-9_]*)\([^)]*\)\s*\{(?:.|\n)*?^}`)
	nameRe := regexp.MustCompile(`function\s+([a-zA-Z_][a-zA-Z0-9_]*)`)
	for _, c := range code {
		matches := re.FindAllString(c, -1)
		for _, match := range matches {
			nameMatch := nameRe.FindStringSubmatch(match)
			if len(nameMatch) > 1 {
				functionName := nameMatch[1]
				data[comName][functionName] = match
			}
		}
	}
}

func Get(comName, funName string) string {
	if data[comName] == nil {
		return ""
	}
	return data[comName][funName]
}

func GetAll(comName string) map[string]string {
	return data[comName]
}

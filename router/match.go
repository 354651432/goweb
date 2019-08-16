package router

import (
	regexp2 "regexp"
	"strings"
)

func match(method, path, exp string) bool {
	method = strings.ToLower(method)

	idx := strings.Index(exp, ":")
	if idx == -1 {
		return false
	}

	method1 := exp[0:idx]
	if method1 != "any" && method != method1 {
		return false
	}

	exp1 := exp[idx+1:]
	if !strings.Contains(exp1, "/") {
		exp1 = "/" + exp1
	}

	rules := strings.Split(exp1, "/")
	pathSplit := strings.Split(path, "/")
	lenPathSplit := len(pathSplit)
	rulesLen := len(rules)

	id := 1
	for ; id < rulesLen; id++ {
		if id >= lenPathSplit {
			return false
		}
		rule := rules[id]

		if !strings.Contains(rule, "@") && rule != pathSplit[id] {
			return false
		}

		if !strings.Contains(rule, ":") {
			params[strings.Trim(rule, "@")] = pathSplit[id]
			continue
		}

		arr := strings.Split(rule, ":")
		name := arr[0]
		regexp := arr[1]

		if matched, err := regexp2.Match(regexp, []byte(pathSplit[id])); matched && err == nil {
			params[strings.Trim(name, "@")] = pathSplit[id]
			continue
		}
		return false
	}

	return id == lenPathSplit
}

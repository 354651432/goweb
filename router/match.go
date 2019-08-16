package router

import (
	"fmt"
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
	if method != method1 {
		return false
	}

	if !strings.Contains(exp[idx+1:], "/") {
		return false
	}

	rules := strings.Split(exp[idx+1:], "/")
	pathSplit := strings.Split(path, "/")
	lenPathSplit := len(pathSplit)
	rulesLen := len(rules)

	for id := 1; id < rulesLen; id++ {
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

		var name, regexp string
		_, err := fmt.Sscanf(rule, "%v:%v", &name, &regexp)
		if err != nil {

		}

		if matched, err := regexp2.Match(regexp, []byte(pathSplit[id])); matched && err == nil {
			params[strings.Trim(rule, "@")] = pathSplit[id]
			continue
		}
		return false
	}

	return true
}

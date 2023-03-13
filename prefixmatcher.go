package utils

import (
	"os"
	"strings"
)

type PrefixMatcher map[string]interface{}

func (p *PrefixMatcher) Add(str string) {
	paths := strings.Split(str, string(os.PathSeparator))
	p.AddNode(paths)
}

func (p *PrefixMatcher) Match(str string) bool {
	paths := strings.Split(str, string(os.PathSeparator))
	return p.MatchNode(paths)
}

func (p *PrefixMatcher) AddNode(path []string) {
	if len(path) == 0 {
		return
	}
	cur, next := path[0], path[1:]
	child, exist := (*p)[cur]
	if exist {
		child := child.(PrefixMatcher)
		if len(child) == 0 {
			return
		} else {
			child.AddNode(next)
		}
	} else {
		var newChild PrefixMatcher = make(map[string]interface{})
		(*p)[cur] = newChild
		newChild.AddNode(next)
	}
}

func (p *PrefixMatcher) MatchNode(path []string) bool {
	if len(path) == 0 {
		return false
	}
	child, exist := (*p)[path[0]]
	if exist {
		child := child.(PrefixMatcher)
		if len(child) == 0 {
			return true
		} else {
			child.MatchNode(path[1:])
		}
	}
	return false
}

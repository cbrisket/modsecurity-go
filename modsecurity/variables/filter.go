package variables

import (
	"regexp"

	"github.com/senghoo/modsecurity-go/utils"
)

type filter struct {
	include      []string
	includeRegex []*regexp.Regexp
	exclude      []string
	excludeRegex []*regexp.Regexp
}

func isRegex(str string) bool {
	if len(str) < 2 {
		return false
	}
	return str[0] == '/' && str[len(str)-1] == '/'
}

func (f *filter) Include(s string) error {
	if isRegex(s) {
		re, err := regexp.Compile(s[1 : len(s)-1])
		if err != nil {
			return err
		}
		f.includeRegex = append(f.includeRegex, re)
		return nil
	}
	f.include = append(f.include, s)
	return nil
}
func (f *filter) Exclude(s string) error {
	if isRegex(s) {
		re, err := regexp.Compile(s[1 : len(s)-1])
		if err != nil {
			return err
		}
		f.excludeRegex = append(f.excludeRegex, re)
		return nil
	}
	f.exclude = append(f.exclude, s)
	return nil
}

func (f *filter) Fetch(vs map[string][]string) []string {
	var res []string
	withName := f.FetchWithNames(vs)
	for _, v := range withName {
		res = append(res, v...)
	}
	return res
}

func (f *filter) Names(vs map[string][]string) []string {
	var res []string
	withName := f.FetchWithNames(vs)
	for k, _ := range withName {
		res = append(res, k)
	}
	return res
}

func (f *filter) FetchWithNames(vs map[string][]string) map[string][]string {
	var res = make(map[string][]string)
	for k, v := range vs {
		if (f.include != nil && !utils.StringInSlice(k, f.include)) ||
			(f.includeRegex != nil && !utils.StringInRegexSlice(k, f.includeRegex)) {
			continue
		}
		if (f.exclude != nil && utils.StringInSlice(k, f.exclude)) ||
			(f.excludeRegex != nil && utils.StringInRegexSlice(k, f.excludeRegex)) {
			continue
		}
		res[k] = append(res[k], v...)
	}
	return res
}

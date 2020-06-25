package inject

import "strings"

func simpleReplace(target []string, replace, with string) []string {
	out := make([]string, len(target))
	for i := range target {
		out[i] = strings.ReplaceAll(target[i], replace, with)
	}

	return out
}

func forceQuotesReplace(target []string, replace, with string) []string {
	out := make([]string, len(target))

	for i := range target {
		if strings.Index(target[i], `"`+replace+`"`) == -1 {
			out[i] = strings.ReplaceAll(target[i], replace, `"`+with+`"`)
		} else {
			out[i] = strings.ReplaceAll(target[i], replace, with)
		}
	}

	return out
}

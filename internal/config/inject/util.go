package inject

import "strings"

func simpleReplace(target []string, replace, with string) []string {
	out := make([]string, len(target))
	for i := range target {
		out[i] = strings.ReplaceAll(target[i], replace, with)
	}

	return out
}

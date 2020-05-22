package inject

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

var outputFileInjectorTarget = regexp.MustCompile(`"?<<output-files(?:\[([^]]*)](?:\[([^]]*)])?)?>>"?`)

const (
	simpleOutFileTarget  = "<<output-files>>"
	wrappedOutFileTarget = `"<<output-files>>"`
)

func NewOutputFileInjector(details *job.Details) VariableInjector {
	return &outputFileInjector{details}
}

type outputFileInjector struct {
	state *job.Details
}

func (t *outputFileInjector) Inject(targets []string) ([]string, error) {
	out := make([]string, 0, len(targets))
	for _, target := range targets {
		matches := outputFileInjectorTarget.FindAllStringSubmatchIndex(target, -1)

		// If the pattern doesn't exist, don't process the string
		if matches == nil {
			out = append(out, target)
			continue
		}

		for _, match := range matches {
			fullMatch := target[match[0]:match[1]]

			// Unwrapped, space separated string
			if fullMatch == simpleOutFileTarget {
				out = t.simpleAll(out, target, match)
				continue
			}

			// Quote wrapped, space separated string
			if fullMatch == wrappedOutFileTarget {
				out = t.wrappedAll(out, target, match)
				continue
			}

			// 2D array
			if match[4] > -1 {
				tmp, err := t.handle2dArray(out, target, match)
				if err != nil {
					return nil, err
				}
				out = tmp
				continue
			}

			// 1D array
			if tmp, err := t.handle1dArray(out, target, match); err != nil {
				return nil, err
			} else {
				out = tmp
			}
		}
	}

	return out, nil
}

func (t *outputFileInjector) simpleAll(out []string, target string, match []int) []string {
	if match[0] > 0 {
		out = append(out, target[:match[0]])
	}

	for _, file := range t.state.OutputFiles[len(t.state.OutputFiles)-1] {
		out = append(out, file)
	}

	if match[1] < len(target) {
		out = append(out, target[match[1]:])
	}

	return out
}

func (t *outputFileInjector) wrappedAll(out []string, target string, match []int) []string {
	return append(out, target[:match[0]+1]+
		strings.Join(t.state.OutputFiles[len(t.state.OutputFiles)-1], " ")+
		target[match[1]-1:])
}

func (t *outputFileInjector) handle2dArray(
	out []string,
	target string,
	match []int,
) ([]string, error) {
	x, err := strconv.Atoi(target[match[2]:match[3]])
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(target[match[4]:match[5]])
	if err != nil {
		return nil, err
	}

	if x >= len(t.state.OutputFiles) {
		return nil, fmt.Errorf("invalid command index %d, array size is %d",
			x, len(t.state.OutputFiles))
	}

	if y >= len(t.state.OutputFiles[x]) {
		return nil, fmt.Errorf(
			"invalid output file index %d on command %d, array size is %d",
			y, x, len(t.state.OutputFiles[x]))
	}

	// if wrapped
	if target[match[0]] == '"' && target[match[1]-1] == '"' {
		return append(out, target[:match[0]+1]+t.state.OutputFiles[x][y]+
			target[match[1]-1:]), nil
	}

	// if unwrapped
	return append(out, target[:match[0]]+t.state.OutputFiles[x][y]+
		target[match[1]:]), nil
}

func (t *outputFileInjector) handle1dArray(out []string, target string, match []int) ([]string, error) {
	x, err := strconv.Atoi(target[match[2]:match[3]])

	// TODO: improve this error
	if err != nil {
		return nil, err
	}

	if x >= len(t.state.OutputFiles) {
		// TODO: improve this error
		return nil, fmt.Errorf("invalid command index %d, array size is %d",
			x, len(t.state.OutputFiles))
	}

	if target[match[0]] == '"' && target[match[1]-1] == '"' {
		return append(out, target[:match[0]+1]+
			strings.Join(t.state.OutputFiles[x], " ")+target[match[1]-1:]), nil
	}

	if match[0] > 0 {
		out = append(out, target[:match[0]])
	}

	for _, file := range t.state.OutputFiles[x] {
		out = append(out, file)
	}

	if match[1] < len(target) {
		out = append(out, target[match[1]:])
	}

	return out, nil
}

package inject

import (
	"errors"
	"fmt"
	"github.com/VEuPathDB/util-exporter-server/internal/process"
	"regexp"
	"strconv"
	"strings"
)

var inputFileInjectorTarget = regexp.MustCompile(`"?<<input-files(?:\[(\d*)])?>>"?`)

const (
	simpleInputFileTarget     = "<<input-files>>"
	allInputFileTarget        = "<<input-files[]>>"
	wrappedInputFileTarget    = `"<<input-files>>"`
	wrappedAllInputFileTarget = `"<<input-files[]>>"`
)

func NewInputFileInjector(details *process.Details) VariableInjector {
	return &timeInjector{details}
}

type inputFileInjector struct {
	state *process.Details
}

func (t *inputFileInjector) Inject(target []string) ([]string, error) {
	out := make([]string, 0, len(target))
	for _, tgt := range target {
		matches := inputFileInjectorTarget.FindAllStringSubmatchIndex(tgt, -1)

		// If the pattern doesn't exist, don't process the string
		if matches == nil {
			out = append(out, tgt)
			continue
		}

		// TODO: Need to replace the token
		for _, match := range matches {
			switch tgt[match[0]:match[1]] {
			case simpleInputFileTarget, allInputFileTarget:
				out = t.simpleAll(out, tgt, match)
				continue

			case wrappedInputFileTarget, wrappedAllInputFileTarget:
				out = t.wrappedAll(out, tgt, match)
				continue
			}

			// this will catch empty number value as well as no match
			if match[2] == match[3] {
				// TODO: improve this error
				return nil, errors.New("invalid state: input file variable injector")
			}

			if tmp, err := t.targetFile(out, tgt, match); err != nil {
				return nil, err
			} else {
				out = tmp
			}
		}
	}
}

func (t *inputFileInjector) simpleAll(out []string, target string, match []int) []string {
	if match[0] > 0 {
		out = append(out, target[:match[0]])
	}

	for _, file := range t.state.InputFiles {
		out = append(out, file)
	}

	if match[1] < len(target) {
		out = append(out, target[match[1]:])
	}

	return out
}

func (t *inputFileInjector) wrappedAll(out []string, target string, match []int) []string {
	return append(out, target[:match[0]+1] +
		strings.Join(t.state.InputFiles, " ") +
		target[match[1]-1:])
}

func (t *inputFileInjector) targetFile(out []string, target string, match []int) ([]string, error) {
	index, err := strconv.Atoi(target[match[2]:match[3]])

	// TODO: improve this error
	if err != nil {
		return nil, err
	}

	if index >= len(t.state.InputFiles) {
		// TODO: improve this error
		return nil, fmt.Errorf("invalid input file index %d, array size is %d",
			index, len(t.state.InputFiles))
	}

	return append(out, target[:match[0]] + t.state.InputFiles[index] +
		target[match[1]:]), nil
}
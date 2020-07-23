package inject

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

var inputFileInjectorTarget = regexp.MustCompile(`"?<<input-files(?:\[([^]]*)])?>>"?`)

const (
	simpleInputFileTarget  = "<<input-files>>"
	wrappedInputFileTarget = `"<<input-files>>"`
)

// NewInputFileInjector returns a new VariableInjector instance that will
// replace <<input-files>> and <<input-files[n]>> variables in a command config.
func NewInputFileInjector(
	details *job.Details,
	_ *job.Metadata,
	log *logrus.Entry,
) VariableInjector {
	log.Trace("inject.NewInputFileInjector")
	return &inputFileInjector{log, details}
}

type inputFileInjector struct {
	log   *logrus.Entry
	state *job.Details
}

func (t *inputFileInjector) Inject(target []string) ([]string, error) {
	t.log.Trace("inject.inputFileInjector.Inject")
	out := make([]string, 0, len(target))

	for _, tgt := range target {
		matches := inputFileInjectorTarget.FindAllStringSubmatchIndex(tgt, -1)

		// If the pattern doesn't exist, don't process the string
		if matches == nil {
			out = append(out, tgt)
			continue
		}

		for _, match := range matches {
			switch tgt[match[0]:match[1]] {
			case simpleInputFileTarget:
				out = t.simpleAll(out, tgt, match)
				continue

			case wrappedInputFileTarget:
				out = t.wrappedAll(out, tgt, match)
				continue
			}

			// this will catch empty number value as well as no match
			if match[2] == match[3] {
				// TODO: improve this error
				return nil, errors.New("invalid state: input file variable injector")
			}

			tmp, err := t.targetFile(out, tgt, match)
			if err != nil {
				return nil, err
			}

			out = tmp
		}
	}

	return out, nil
}

func (t *inputFileInjector) simpleAll(out []string, target string, match []int) []string {
	t.log.Trace("inject.inputFileInjector.simpleAll")
	if match[0] > 0 {
		out = append(out, target[:match[0]])
	}

	out = append(out, t.state.UnpackedFiles...)

	if match[1] < len(target) {
		out = append(out, target[match[1]:])
	}

	return out
}

func (t *inputFileInjector) wrappedAll(out []string, target string, match []int) []string {
	t.log.Trace("inject.inputFileInjector.wrappedAll")
	return append(out, target[:match[0]+1]+
		strings.Join(t.state.UnpackedFiles, " ")+
		target[match[1]-1:])
}

func (t *inputFileInjector) targetFile(out []string, target string, match []int) ([]string, error) {
	t.log.Trace("inject.inputFileInjector.targetFile")

	index, err := strconv.Atoi(target[match[2]:match[3]])

	// TODO: improve this error
	if err != nil {
		return nil, err
	}

	if index >= len(t.state.UnpackedFiles) {
		// TODO: improve this error
		return nil, fmt.Errorf("invalid input file index %d, array size is %d",
			index, len(t.state.UnpackedFiles))
	}

	if target[match[0]] == '"' && target[match[1]-1] == '"' {
		return append(out, target[:match[0]+1]+t.state.UnpackedFiles[index]+
			target[match[1]-1:]), nil
	}

	return append(out, target[:match[0]]+t.state.UnpackedFiles[index]+
		target[match[1]:]), nil
}

package inject

import (
	"errors"
	"regexp"

	"github.com/sirupsen/logrus"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

var handlerParamInjectorTarget = regexp.MustCompile(`"?<<handler-params\.([a-z0-9A-Z_-]+)>>"?`)

// HandlerParamInjector returns a new VariableInjector instance that will
// replace <<handler-params.X>> variables in a command config.
func NewHandlerParamInjector(
	details *job.Details,
	metadata *job.Metadata,
	log *logrus.Entry,
) VariableInjector {
	log.Trace("inject.NewHandlerParamInjector")
	return &handlerParamInjector{log, metadata}
}

type handlerParamInjector struct {
	log      *logrus.Entry
	metadata *job.Metadata
}

func (t *handlerParamInjector) Inject(target []string) ([]string, error) {
	t.log.Trace("inject.handlerParamInjector.Inject")
	out := make([]string, 0, len(target))

	for _, tgt := range target {
		matches := handlerParamInjectorTarget.FindAllStringSubmatchIndex(tgt, -1)

		// If the pattern doesn't exist, don't process the string
		if matches == nil {
			out = append(out, tgt)
			continue
		}

		for _, match := range matches {
			if t.metadata.HandlerParams == nil {
				return nil, errors.New("invalid state: HandlerParams are not set on job metadata.")
			}
			handler_param_name := tgt[match[2]:match[3]]
			handler_param_value, is_present := t.metadata.HandlerParams[handler_param_name]
			if !is_present {
				return nil, errors.New("invalid state: Parameter not found in input handler parameters")
			}
			out = append(out, tgt[:match[0]]+handler_param_value+tgt[match[1]:])
		}
	}

	return out, nil
}

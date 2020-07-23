package config

import (
	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/sirupsen/logrus"
)

// InjectorProvider defines a function that is used to construct and populate a
// VariableInjector instance with job context data.
type InjectorProvider func(
	*job.Details,
	*job.Metadata,
	*logrus.Entry,
) inject.VariableInjector

// InjectorList returns a slice of providers for all VariableInjectors.
func InjectorList() []InjectorProvider {
	return []InjectorProvider{
		inject.NewCwdInjector,
		inject.NewDateInjector,
		inject.NewDateTimeInjector,
		inject.NewDsDescriptionInjector,
		inject.NewDsNameInjector,
		inject.NewDsSummaryInjector,
		inject.NewDsUserIdInjector,
		inject.NewDsOriginInjector,
		inject.NewInputFileInjector,
		inject.NewTimeInjector,
		inject.NewTimestampInjector,
	}
}

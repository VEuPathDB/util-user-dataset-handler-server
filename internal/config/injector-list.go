package config

import (
	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

type InjectorProvider func(*job.Details, *job.Metadata) inject.VariableInjector

func InjectorList() []InjectorProvider {
	return []InjectorProvider{
		inject.NewCwdInjector,
		inject.NewDateInjector,
		inject.NewDateTimeInjector,
		inject.NewDsDescriptionInjector,
		inject.NewDsNameInjector,
		inject.NewDsSummaryInjector,
		inject.NewDsUserEmailInjector,
		inject.NewInputFileInjector,
		inject.NewOutputFileInjector,
		inject.NewTimeInjector,
		inject.NewTimestampInjector,
	}
}

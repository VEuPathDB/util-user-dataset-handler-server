package config

import (
	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/process"
)

type InjectorProvider func(*process.Details) inject.VariableInjector

func InjectorList() []InjectorProvider {
	return []InjectorProvider{
		inject.NewCwdInjector,
		inject.NewDateInjector,
		inject.NewDateTimeInjector,
		inject.NewInputFileInjector,
		inject.NewOutputFileInjector,
		inject.NewTimeInjector,
		inject.NewTimestampInjector,
	}
}

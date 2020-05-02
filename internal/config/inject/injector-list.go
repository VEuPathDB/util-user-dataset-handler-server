package inject

import "github.com/VEuPathDB/util-exporter-server/internal/process"

type InjectorProvider func(*process.Details) VariableInjector

func InjectorList() []InjectorProvider {
	return []InjectorProvider{
		NewCwdInjector,
		NewDateInjector,
		NewDateTimeInjector,
		NewInputFileInjector,
		NewOutputFileInjector,
		NewTimeInjector,
		NewTimestampInjector,
	}
}

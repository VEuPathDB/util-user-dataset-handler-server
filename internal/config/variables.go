package config

import (
	"github.com/VEuPathDB/util-exporter-server/internal/process"
	"strconv"
	"time"
)

type VariableProcessor func(*process.Details) ([]string, error)

func VariableProcessors() map[string]VariableProcessor {
	return map[string]VariableProcessor{
		"<<date>>":      varDate,
		"<<time>>":      varTime,
		"<<timestamp>>": varTimestamp,
		"<<date-time>>": varDateTime,
		"<<input-files>>":
	}
}

func var____(dets *process.Details) ([]string, error) {}


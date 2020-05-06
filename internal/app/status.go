package app

const (
	StatusSuccess = 0

	// Used by logrus Fatal
	StatusLoggerFatal = 1

	// Used by config validation run mode
	StatusValidateConfFailed = 2

	// Used by config generation run mode
	StatusGenConfFailed = 3
)

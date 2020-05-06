package app

var Keys = globalKeys{
	Logger: loggerKeys{
		Service: "service",
		Status:  "status",
	},
}

type globalKeys struct {
	Logger loggerKeys
}

type loggerKeys struct {
	Service string
	Status  string
}

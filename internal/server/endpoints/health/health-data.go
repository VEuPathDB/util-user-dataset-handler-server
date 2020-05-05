package health

type Data struct {
	Status  string      `json:"status"`
	Version string      `json:"version"`
	Details interface{} `json:"stats"`
}

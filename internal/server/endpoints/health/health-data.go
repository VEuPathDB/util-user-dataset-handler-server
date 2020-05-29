package health

// Data defines a health response body.
type Data struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

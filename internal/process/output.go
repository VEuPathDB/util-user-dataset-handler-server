package process

type StatusFile struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Files   []string `json:"files"`
	Pack    []string `json:"pack"`
}

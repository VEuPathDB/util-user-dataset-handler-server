package script

type ResponseStatus string

const (
	StatusSuccess ResponseStatus = "success"
	StatusError   ResponseStatus = "error"
)

type Response struct {
	Status   ResponseStatus `json:"status"`
	Message  string         `json:"message"`
	Warnings []string       `json:"warnings"`
	Files    []string       `json:"files"`
}

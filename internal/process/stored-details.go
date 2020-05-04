package process

import "time"

// StorableDetails is the subset of process details that
// will be kept in memory for up to 3 days.
type StorableDetails struct {
	// Started is the timestamp from the start of the HTTP
	// request.
	Started time.Time `json:"started"`

	// UserID is the WDK user ID for the user initiating the
	// request.
	UserID uint `json:"userId"`

	// Token is the unique ID given to this request for
	// tracking purposes.
	Token string `json:"token"`

	// Status is the execution status for the current request.
	Status Status `json:"status"`

	Size uint `json:"size"`
}

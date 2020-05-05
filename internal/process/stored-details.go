package process

import (
	"github.com/VEuPathDB/util-exporter-server/internal/wdk/site"
	"time"
)

// StorableDetails is the subset of process details that
// will be kept in memory for up to 3 days.
type StorableDetails struct {
	// Started is the timestamp from the start of the HTTP
	// request.
	Started time.Time `json:"started"`

	// Duration is the length of time the request took
	// overall.
	Duration time.Duration `json:"duration"`

	// UserID is the WDK user ID for the user initiating the
	// request.
	UserID uint `json:"userId"`

	// Token is the unique ID given to this request for
	// tracking purposes.
	Token string `json:"token"`

	// Status is the execution status for the current request.
	Status Status `json:"status"`

	// Projects holds the target projects for the request.
	Projects []site.WdkSite `json:"projects"`

	// Size contains the unpackaged size of the dataset
	// payload.
	Size uint `json:"size"`
}

package job

// Status defines a job status value
type Status string

// All possible job status values.
const (
	StatusReceiving  Status = "uploading files for processing"
	StatusUnpacking  Status = "unpacking uploaded files"
	StatusProcessing Status = "processing uploaded files"
	StatusCompleted  Status = "job completed"
	StatusFailed     Status = "job failed"
	StatusNotStarted Status = "awaiting upload"
)

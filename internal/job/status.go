package job

type Status string

const (
	StatusReceiving  Status = "uploading files for processing"
	StatusUnpacking  Status = "unpacking uploaded files"
	StatusProcessing Status = "processing uploaded files"
	StatusCompleted  Status = "job completed"
	StatusFailed     Status = "job failed"
	StatusNotStarted Status = "awaiting upload"
)

package process

type Status string

const (
	StatusReceiving  Status = "uploading files for processing"
	StatusUnpacking  Status = "unpacking uploaded files"
	StatusProcessing Status = "processing uploaded files"
	StatusPacking    Status = "packing processed files"
	StatusSending    Status = "downloading processed files"
	StatusCompleted  Status = "job completed"
	StatusFailed     Status = "job failed"
	StatusNotStarted Status = "awaiting upload"
)

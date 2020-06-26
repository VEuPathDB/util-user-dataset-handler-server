package job

type Details struct {
	StorableDetails

	// InputFile is the uploaded tar file name.
	InputFile string `json:"uploadName"`

	// UnpackedFiles is the list of unpacked files uploaded for
	// processing.
	UnpackedFiles []string `json:"files"`

	// OutputFiles is a two-dimensional array of command
	// executions and their output files.
	OutputFiles [][]string `json:"outputFiles"`

	// WorkingDir is the full path to the work directory
	// created for this request.
	WorkingDir string `json:"workingDir"`
}

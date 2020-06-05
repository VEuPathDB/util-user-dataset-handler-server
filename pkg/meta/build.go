package meta

import "fmt"

const buildRenderFormat =
`Version:    %s
Commit:     %s
Build Date: %s`

// Values populated automatically at build time.
var (
	version   = "untagged dev build"
	commit    = "unknown"
	buildDate = "unknown"
)

func GetBuildMeta() BuildMeta {
	return BuildMeta{
		Version:   version,
		Commit:    commit,
		BuildDate: buildDate,
	}
}

type BuildMeta struct {
	Version   string
	Commit    string
	BuildDate string
}

func (b BuildMeta) String() string {
	return fmt.Sprintf(buildRenderFormat, b.Version, b.Commit, b.BuildDate)
}
package site

type WdkSite string

// Enum of allowed ProjectType values.
const (
	AmoebaDB        WdkSite = "AmoebaDB"
	ClinEpiDB       WdkSite = "ClinEpiDB"
	CryptoDB        WdkSite = "CryptoDB"
	FungiDB         WdkSite = "FungiDB"
	GiardiaDB       WdkSite = "GiardiaDB"
	HostDB          WdkSite = "HostDB"
	MicrobiomeDB    WdkSite = "MicrobiomeDB"
	MicrosporidiaDB WdkSite = "MicrosporidiaDB"
	OrthoMCL        WdkSite = "OrthoMCL"
	PiroplasmaDB    WdkSite = "PiroplasmaDB"
	PlasmoDB        WdkSite = "PlasmoDB"
	SchistoDB       WdkSite = "SchistoDB"
	ToxoDB          WdkSite = "ToxoDB"
	TrichDB         WdkSite = "TrichDB"
	TritrypDB       WdkSite = "TritrypDB"
	VectorBase      WdkSite = "VectorBase"
	VEuPathDB       WdkSite = "VEuPathDB"
)

var validSites = map[WdkSite]bool{
	AmoebaDB:        true,
	ClinEpiDB:       true,
	CryptoDB:        true,
	FungiDB:         true,
	GiardiaDB:       true,
	HostDB:          true,
	MicrobiomeDB:    true,
	MicrosporidiaDB: true,
	OrthoMCL:        true,
	PiroplasmaDB:    true,
	PlasmoDB:        true,
	SchistoDB:       true,
	ToxoDB:          true,
	TrichDB:         true,
	TritrypDB:       true,
	VectorBase:      true,
	VEuPathDB:       true,
}

// IsValid returns whether or not the the wrapped value is a known WdkSite.
func (w WdkSite) IsValid() bool {
	return validSites[w]
}

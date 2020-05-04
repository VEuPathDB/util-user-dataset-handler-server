package site

type WdkSite string

const (
	AmoebaDB        WdkSite = "AmoebaDB"
	ClinepiDB       WdkSite = "ClinepiDB"
	CryptoDB        WdkSite = "CryptoDB"
	FungiDB         WdkSite = "FungiDB"
	GiardiaDB       WdkSite = "GiardiaDB"
	HostDB          WdkSite = "HostDB"
	MicrobiomeDB    WdkSite = "MicrobiomeDB"
	MicrosporidiaDB WdkSite = "MicrosporidiaDB"
	Orthomcl        WdkSite = "Orthomcl"
	PiroplasmaDB    WdkSite = "PiroplasmaDB"
	PlasmoDB        WdkSite = "PlasmoDB"
	SchistoDB       WdkSite = "SchistoDB"
	ToxoDB          WdkSite = "ToxoDB"
	TrichDB         WdkSite = "TrichDB"
	TritrypDB       WdkSite = "TritrypDB"
	Vectorbase      WdkSite = "Vectorbase"
	VeupathDB       WdkSite = "VeupathDB"
)

var validSites = map[WdkSite]bool{
	AmoebaDB:        true,
	ClinepiDB:       true,
	CryptoDB:        true,
	FungiDB:         true,
	GiardiaDB:       true,
	HostDB:          true,
	MicrobiomeDB:    true,
	MicrosporidiaDB: true,
	Orthomcl:        true,
	PiroplasmaDB:    true,
	PlasmoDB:        true,
	SchistoDB:       true,
	ToxoDB:          true,
	TrichDB:         true,
	TritrypDB:       true,
	Vectorbase:      true,
	VeupathDB:       true,
}

func (W WdkSite) IsValid() bool {
	return validSites[W]
}

package data

type ClientVersion struct {
	Version string `json:"version"`
	API string `json:"api"`
}

type ClientVersions struct {
	Latest string `json:"latest"`
	Versions []ClientVersion `json:"versions"`
}
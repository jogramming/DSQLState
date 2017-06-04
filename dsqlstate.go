package dsqlstate

import (
	"strconv"
)

const (
	VersionMajor = 0
	VersionMinor = 1
	VersionPatch = 0
)

var (
	VersionString = strconv.Itoa(VersionMajor) + "." + strconv.Itoa(VersionMinor) + "." + strconv.Itoa(VersionPatch)
)

package dsqlstate

import (
	"strconv"
)

const (
	VersionMajor = 0
	VersionMinor = 2
	VersionPatch = 1
)

var (
	VersionString = strconv.Itoa(VersionMajor) + "." + strconv.Itoa(VersionMinor) + "." + strconv.Itoa(VersionPatch)
)

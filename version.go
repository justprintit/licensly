package licensly

import _ "embed"

//go:embed version.txt
var version string

func Version() string {
	return version
}

package version

import (
	"strconv"
	"time"
)

var (
	// BuildTime is the build time of the binary (set externally with ldflags).
	BuildTime = "unknown"

	// CommitHash is the Git hash of the branch, used for version purposes (set externally with ldflags).
	CommitHash = "undefined"
)

// BuildTimeFormatted method returns the build time preserving the RFC3339 format.
func BuildTimeFormatted() string {
	if BuildTime == "unknown" {
		return BuildTime
	}

	seconds, err := strconv.ParseInt(BuildTime, 10, 64)
	if err != nil {
		return "invalid"
	}
	return time.Unix(seconds, 0).Format(time.RFC3339)
}
